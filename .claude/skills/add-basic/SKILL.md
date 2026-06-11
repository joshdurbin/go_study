---
name: add-basic
description: Add a new Basics lesson to the Go study guide — a worked example of a Go language feature (syntax, stdlib idiom, concurrency primitive). Use when the user wants to add a new topic under basics/. Creates one .go file (the lesson source) plus one .md file (study notes).
---

# Add a basic lesson

Creates a new lesson under `basics/`. Each basic is a self-contained worked example that compiles, runs, and demonstrates ONE coherent concept.

## Files to create

For a new topic named `<topic>` (snake_case, e.g., `panic_recover_idioms`):

| File | Purpose |
|------|---------|
| `basics/NN_<topic>.go`  | Worked example. `package main` with `func main()`. **First line must be `//go:build ignore`** (excludes from package builds; `go run` still works). |
| `basics/NN_<topic>.md`  | Study notes rendered as markdown above the editor. Tight, ~5-15 short lines. |

`NN` is the next available numeric prefix (currently lessons 01–22; new lesson is 23 unless inserting). The prefix drives sidebar order, not categorization — pick the lowest unused number.

## How to choose the topic

A good basic teaches ONE concept:
- A language feature (`select`, `range over int`, generics constraints)
- A stdlib idiom (`io.Copy`, `errgroup`, `context.WithTimeout`)
- A common gotcha (string vs. []byte, nil interface)

Avoid:
- Topics already covered (read existing `basics/*.md` first; especially overlapping ones are flagged in CLAUDE.md)
- Pattern-level material (that goes in `patterns/`)

## .go file template

```go
//go:build ignore

package main

import "fmt"

// Short comment about what this demonstrates.
// Inline comments explain non-obvious lines.

func main() {
    // self-contained example with printed output
    fmt.Println("...")
}
```

Style rules:
- One `main()` per file. No subcommands or test harness.
- Comments explain WHY, not WHAT. Don't restate code.
- Output should be deterministic and short. Long output dilutes the lesson.
- If demonstrating multiple sub-concepts in one file, separate them with `// ─── Concept name ─────` banners.

## .md file template

```markdown
# <Topic Title>

<1-2 sentences: what this teaches and when you'd reach for it.>

## Worth knowing

- <bullet — a specific fact or rule>
- <bullet — a non-obvious behavior>
- <bullet — interaction with another feature, if any>

## Common gotcha

<1-2 sentences about the most common bug or misconception.>
```

Optional sections (only when they add value):
- `## When to use` / `## When NOT to use`
- `## Interview frame` (what question this answers)

Cross-reference related material with relative paths: `See [patterns/16_fan_in_fan_out.md](../patterns/16_fan_in_fan_out.md)`.

## After creating files

1. Verify it compiles and runs: `go run basics/NN_<topic>.go`
2. Verify whole-module build: `go build ./...`
3. The web app picks up the new lesson on next page refresh — no restart needed (dynamic scanning).

## Reference

See `basics/13_goroutines.go` + `basics/13_goroutines.md` for an exemplar lesson. CLAUDE.md has the broader file conventions.
