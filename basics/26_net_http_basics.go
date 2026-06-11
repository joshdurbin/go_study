//go:build ignore

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

// net/http gives you a production HTTP server and client in the stdlib.
// We use httptest.NewServer so this demo runs without binding a real port
// or blocking on ListenAndServe.

func main() {
	// ─── Define a handler ───
	// In a real binary you'd use http.HandleFunc on the default mux, then
	// http.ListenAndServe(":8080", nil) — but that blocks forever.
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "world"
		}
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "hello, %s", name)
	})
	mux.HandleFunc("/echo-header", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "X-Demo=%s", r.Header.Get("X-Demo"))
	})

	// ─── Spin up an ephemeral test server ───
	srv := httptest.NewServer(mux)
	defer srv.Close()

	// ─── Client: simple GET ───
	resp, err := http.Get(srv.URL + "/hello?name=Ada")
	if err != nil {
		panic(err)
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close() // ALWAYS close — leaks goroutines + connections
	fmt.Printf("GET /hello -> %d %q\n", resp.StatusCode, body)

	// ─── Client: build a request with custom headers ───
	// http.NewRequest + Client.Do gives you control over method, headers, body.
	req, _ := http.NewRequest(http.MethodGet, srv.URL+"/echo-header", nil)
	req.Header.Set("X-Demo", "abc")
	resp2, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close() // defer is fine when there's no loop
	body2, _ := io.ReadAll(resp2.Body)
	fmt.Printf("GET /echo-header -> %d %q\n", resp2.StatusCode, body2)
}
