package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/joshdurbin/go_study/internal/analytics"
	"github.com/joshdurbin/go_study/internal/db"
	"github.com/joshdurbin/go_study/internal/db/store"
	"github.com/spf13/viper"
)

//go:embed web/*.html
var webFS embed.FS

type Item struct {
	Name     string
	RelPath  string
	Problem  string
	Runnable bool
}

type Subsection struct {
	Name  string
	Items []Item
}

type Section struct {
	Name        string
	Dir         string
	Description string
	Subs        []Subsection
	Items       []Item
}

// sectionDirs are the fixed top-level directories that may contain lessons.
// safePath uses this list (not the dynamically-loaded `sections`) so security
// validation is independent of whether the user just added a new file.
var sectionDirs = []string{"basics", "patterns", "interview_study", "interview_qa"}

type server struct {
	repoRoot string
	tracker  *analytics.Tracker

	mu       sync.RWMutex // guards sections
	sections []Section
}

func runServe() error {
	repoRoot, err := filepath.Abs(viper.GetString("root"))
	if err != nil {
		return err
	}
	conn, err := openDB()
	if err != nil {
		return err
	}
	defer conn.Close()
	if err := db.Migrate(conn); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	s := &server{
		repoRoot: repoRoot,
		tracker:  analytics.New(conn, viper.GetBool("analytics_enabled")),
	}
	// Initial load just for log feedback / fail-fast. Subsequent requests
	// re-scan to pick up files added or removed while the server is running.
	if err := s.loadSections(); err != nil {
		return fmt.Errorf("load sections: %w", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleIndex)
	mux.HandleFunc("/dashboard", s.handleDashboard)
	mux.HandleFunc("/api/file", s.handleFile)
	mux.HandleFunc("/api/run", s.handleRun)
	mux.HandleFunc("/api/hint", s.handleHint)
	mux.HandleFunc("/api/reveal", s.handleReveal)
	mux.HandleFunc("/api/editor-state", s.handleEditorState)
	mux.HandleFunc("/api/stats", s.handleStats)
	mux.HandleFunc("/ws", s.handleWS)

	handler := s.tracker.SessionMiddleware(mux)

	addr := viper.GetString("addr")
	srv := &http.Server{Addr: addr, Handler: handler}

	go func() {
		log.Printf("go_study serving from %s on http://localhost%s (analytics=%v)",
			repoRoot, addr, s.tracker.Enabled())
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Println("shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return srv.Shutdown(ctx)
}

// currentSections returns a snapshot under read lock.
func (s *server) currentSections() []Section {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.sections
}

func (s *server) loadSections() error {
	specs := []struct{ dir, name, desc string }{
		{"basics", "Basics", "Core Go language: syntax, types, concurrency primitives, stdlib idioms."},
		{"patterns", "Design Patterns", "Idiomatic Go patterns you'll be expected to recognize and write."},
		{"interview_study", "Interview Practice", "Algorithm blocks with problem statements and reference solutions."},
		{"interview_qa", "Interview Q&A", "Short-answer Go interview topics — gotchas, idioms, and trick questions."},
	}
	var built []Section
	for _, spec := range specs {
		dirPath := filepath.Join(s.repoRoot, spec.dir)
		info, err := os.Stat(dirPath)
		if err != nil || !info.IsDir() {
			continue
		}
		sec := Section{Name: spec.name, Dir: spec.dir, Description: spec.desc}
		entries, err := os.ReadDir(dirPath)
		if err != nil {
			return err
		}
		hasSubdirs := false
		for _, e := range entries {
			if e.IsDir() {
				hasSubdirs = true
				break
			}
		}
		if hasSubdirs {
			for _, e := range entries {
				if !e.IsDir() {
					continue
				}
				subPath := filepath.Join(dirPath, e.Name())
				subRel := filepath.Join(spec.dir, e.Name())
				subEntries, err := os.ReadDir(subPath)
				if err != nil {
					return err
				}
				nested := false
				for _, se := range subEntries {
					if se.IsDir() {
						nested = true
						break
					}
				}
				if nested {
					for _, se := range subEntries {
						if !se.IsDir() {
							continue
						}
						items, err := collectItems(filepath.Join(subPath, se.Name()), filepath.Join(subRel, se.Name()))
						if err != nil {
							return err
						}
						if len(items) == 0 {
							continue
						}
						sec.Subs = append(sec.Subs, Subsection{
							Name:  prettify(e.Name()) + " — " + prettify(se.Name()),
							Items: items,
						})
					}
				} else {
					items, err := collectItems(subPath, subRel)
					if err != nil {
						return err
					}
					sec.Subs = append(sec.Subs, Subsection{Name: prettify(e.Name()), Items: items})
				}
			}
		} else {
			items, err := collectItems(dirPath, spec.dir)
			if err != nil {
				return err
			}
			sec.Items = items
		}
		built = append(built, sec)
	}
	s.mu.Lock()
	s.sections = built
	s.mu.Unlock()
	return nil
}

func collectItems(dir, relDir string) ([]Item, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var items []Item
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") {
			continue
		}
		if strings.HasSuffix(e.Name(), ".starter.go") {
			continue // starter scaffolds aren't standalone lessons
		}
		rel := filepath.Join(relDir, e.Name())
		item := Item{
			Name:     prettify(strings.TrimSuffix(e.Name(), ".go")),
			RelPath:  filepath.ToSlash(rel),
			Runnable: true,
		}
		base := filepath.Join(dir, strings.TrimSuffix(e.Name(), ".go"))
		if b, err := os.ReadFile(base + ".md"); err == nil {
			item.Problem = string(b)
		} else if b, err := os.ReadFile(base + ".txt"); err == nil {
			item.Problem = string(b)
		}
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].RelPath < items[j].RelPath })
	return items, nil
}

func prettify(name string) string {
	s := name
	s = strings.TrimPrefix(s, "block")
	for len(s) > 0 && (s[0] >= '0' && s[0] <= '9') {
		s = s[1:]
	}
	s = strings.TrimPrefix(s, "_")
	s = strings.ReplaceAll(s, "_", " ")
	if s == "" {
		return name
	}
	parts := strings.Fields(s)
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, " ")
}

func parseHints(s string) []string {
	lines := strings.Split(s, "\n")
	var hints []string
	var cur strings.Builder
	flush := func() {
		t := strings.TrimSpace(cur.String())
		if t != "" {
			hints = append(hints, t)
		}
		cur.Reset()
	}
	for _, ln := range lines {
		if strings.HasPrefix(ln, "## ") {
			flush()
			continue
		}
		cur.WriteString(ln)
		cur.WriteByte('\n')
	}
	flush()
	return hints
}

func (s *server) safePath(rel string) (string, error) {
	clean := filepath.Clean(rel)
	if strings.HasPrefix(clean, "..") || filepath.IsAbs(clean) {
		return "", fmt.Errorf("invalid path")
	}
	abs := filepath.Join(s.repoRoot, clean)
	relCheck, err := filepath.Rel(s.repoRoot, abs)
	if err != nil || strings.HasPrefix(relCheck, "..") {
		return "", fmt.Errorf("path escapes root")
	}
	allowed := false
	for _, d := range sectionDirs {
		if strings.HasPrefix(filepath.ToSlash(clean), d+"/") {
			allowed = true
			break
		}
	}
	if !allowed {
		return "", fmt.Errorf("path not in a known section")
	}
	if !strings.HasSuffix(abs, ".go") {
		return "", fmt.Errorf("only .go files")
	}
	return abs, nil
}

func (s *server) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Re-scan the lesson tree on every page load. Picks up files added or
	// removed while the server is running without requiring a restart.
	if err := s.loadSections(); err != nil {
		log.Printf("reload sections: %v", err)
	}
	tmpl, err := template.ParseFS(webFS, "web/index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := tmpl.Execute(w, s.currentSections()); err != nil {
		log.Println(err)
	}
}

func (s *server) handleFile(w http.ResponseWriter, r *http.Request) {
	rel := r.URL.Query().Get("path")
	abs, err := s.safePath(rel)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	src, err := os.ReadFile(abs)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	base := strings.TrimSuffix(abs, ".go")
	var problem string
	if b, err := os.ReadFile(base + ".md"); err == nil {
		problem = string(b)
	} else if b, err := os.ReadFile(base + ".txt"); err == nil {
		problem = string(b)
	}
	var hints []string
	if b, err := os.ReadFile(base + ".steps.md"); err == nil {
		hints = parseHints(string(b))
	}
	// Starter: prefer a sibling .starter.go (interview problems). If absent,
	// basics/patterns use their full source as the editor starting point;
	// interview problems fall back to a stub.
	starter := ""
	if b, err := os.ReadFile(base + ".starter.go"); err == nil {
		starter = string(b)
	} else if strings.HasPrefix(rel, "interview_study/") {
		starter = "//go:build ignore\n\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n\t// TODO: implement\n\tfmt.Println()\n}\n"
	} else {
		starter = string(src)
	}
	_ = s.tracker.RecordOpen(r.Context(), analytics.SessionID(r), rel)
	saved, _ := s.tracker.Queries().GetEditorState(r.Context(), rel)
	writeJSON(w, map[string]any{
		"path":    rel,
		"source":  string(src),
		"starter": starter,
		"saved":   saved,
		"problem": problem,
		"hints":   hints,
	})
}

func (s *server) handleRun(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", 405)
		return
	}
	var req struct {
		Path  string `json:"path"`
		Code  string `json:"code"`  // editor contents; if empty, run the file at Path
		Stdin string `json:"stdin"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	abs, err := s.safePath(req.Path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	// Decide what to run: user-edited code (write to temp file) or the original.
	target := abs
	if strings.TrimSpace(req.Code) != "" {
		tmp, err := os.CreateTemp("", "gostudy-*.go")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer os.Remove(tmp.Name())
		if _, err := tmp.WriteString(req.Code); err != nil {
			tmp.Close()
			http.Error(w, err.Error(), 500)
			return
		}
		tmp.Close()
		target = tmp.Name()
	}
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()
	start := time.Now()
	cmd := exec.CommandContext(ctx, "go", "run", target)
	cmd.Dir = filepath.Dir(abs)
	if req.Stdin != "" {
		cmd.Stdin = strings.NewReader(req.Stdin)
	}
	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	runErr := cmd.Run()
	elapsed := time.Since(start)
	_ = s.tracker.RecordRun(r.Context(), analytics.SessionID(r), req.Path,
		elapsed.Milliseconds(), runErr == nil)
	resp := map[string]any{
		"stdout":     stdout.String(),
		"stderr":     stderr.String(),
		"elapsed_ms": elapsed.Milliseconds(),
	}
	if runErr != nil {
		resp["error"] = runErr.Error()
		if ctx.Err() == context.DeadlineExceeded {
			resp["error"] = "timeout after 15s"
		}
	}
	writeJSON(w, resp)
}

func (s *server) handleHint(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", 405)
		return
	}
	var req struct {
		Path      string `json:"path"`
		HintIndex int64  `json:"hint_index"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if _, err := s.safePath(req.Path); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	_ = s.tracker.RecordHint(r.Context(), analytics.SessionID(r), req.Path, req.HintIndex)
	w.WriteHeader(204)
}

func (s *server) handleReveal(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", 405)
		return
	}
	var req struct {
		Path string `json:"path"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if _, err := s.safePath(req.Path); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	_ = s.tracker.RecordSolution(r.Context(), analytics.SessionID(r), req.Path)
	w.WriteHeader(204)
}

func (s *server) handleEditorState(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", 405)
		return
	}
	var req struct {
		Path    string `json:"path"`
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if _, err := s.safePath(req.Path); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	_ = s.tracker.Queries().UpsertEditorState(r.Context(), store.UpsertEditorStateParams{
		LessonPath: req.Path,
		Content:    req.Content,
	})
	w.WriteHeader(204)
}

// handleWS upgrades to a websocket. The client sends frames like:
//
//	{"type":"focus", "path":"basics/01_hello_world.go"}
//	{"type":"blur"}
//	{"type":"tick", "seconds": 5}
//
// On "tick" the server attributes `seconds` to the currently-focused lesson.
func (s *server) handleWS(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true, // same-origin only in practice
	})
	if err != nil {
		log.Println("ws accept:", err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "closing")

	sid := analytics.SessionID(r)
	currentPath := ""
	ctx := r.Context()

	for {
		var msg struct {
			Type    string `json:"type"`
			Path    string `json:"path"`
			Seconds int64  `json:"seconds"`
		}
		if err := wsjson.Read(ctx, c, &msg); err != nil {
			break
		}
		switch msg.Type {
		case "focus":
			if _, err := s.safePath(msg.Path); err == nil {
				currentPath = msg.Path
			}
		case "blur":
			currentPath = ""
		case "tick":
			if currentPath == "" || msg.Seconds <= 0 {
				continue
			}
			_ = s.tracker.RecordTime(ctx, sid, currentPath, msg.Seconds)
		}
	}
	c.Close(websocket.StatusNormalClosure, "")
}

// handleStats returns the aggregated analytics for the dashboard.
func (s *server) handleStats(w http.ResponseWriter, r *http.Request) {
	q := s.tracker.Queries()
	overall, _ := q.OverallStats(r.Context())
	sections, _ := q.SectionTotals(r.Context())
	lessons, _ := q.LessonTotals(r.Context())
	hints, _ := q.HintsPerLesson(r.Context())

	writeJSON(w, map[string]any{
		"overall":  overallToMap(overall),
		"sections": sectionsToMaps(sections),
		"lessons":  lessonsToMaps(lessons),
		"hints":    hints,
	})
}

// coerceInt64 handles sqlc's interface{} type for COALESCE columns.
func coerceInt64(v any) int64 {
	switch x := v.(type) {
	case int64:
		return x
	case float64:
		return int64(x)
	case nil:
		return 0
	default:
		return 0
	}
}

func overallToMap(o store.OverallStatsRow) map[string]any {
	return map[string]any{
		"sessions":         o.Sessions,
		"opens":            o.Opens,
		"total_seconds":    coerceInt64(o.TotalSeconds),
		"hint_reveals":     o.HintReveals,
		"solution_reveals": o.SolutionReveals,
		"code_runs":        o.CodeRuns,
		"successful_runs":  o.SuccessfulRuns,
	}
}

func sectionsToMaps(rows []store.SectionTotalsRow) []map[string]any {
	out := make([]map[string]any, 0, len(rows))
	for _, r := range rows {
		out = append(out, map[string]any{
			"section":          r.Section,
			"opens":            r.Opens,
			"total_seconds":    coerceInt64(r.TotalSeconds),
			"hint_reveals":     r.HintReveals,
			"solution_reveals": r.SolutionReveals,
			"code_runs":        r.CodeRuns,
		})
	}
	return out
}

func lessonsToMaps(rows []store.LessonTotalsRow) []map[string]any {
	out := make([]map[string]any, 0, len(rows))
	for _, r := range rows {
		out = append(out, map[string]any{
			"lesson_path":      r.LessonPath,
			"section":          r.Section,
			"opens":            r.Opens,
			"total_seconds":    coerceInt64(r.TotalSeconds),
			"hint_reveals":     r.HintReveals,
			"solution_reveals": r.SolutionReveals,
			"code_runs":        r.CodeRuns,
		})
	}
	return out
}

func (s *server) handleDashboard(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(webFS, "web/dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		log.Println(err)
	}
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
