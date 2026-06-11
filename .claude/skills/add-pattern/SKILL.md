---
name: add-pattern
description: Add a new design pattern to the Go study guide under patterns/ — an idiomatic Go pattern (functional options, errgroup, circuit breaker, etc.) that learners should recognize and write. Use when the user wants to add a new pattern. Creates one .go file (the pattern implementation) plus one .md file (when to use / not to use).
---

# Add a design pattern

Creates a new lesson under `patterns/`. Patterns are reusable, idiomatic shapes — heavier than basics, with the framing of "when do I reach for this?"

## Files to create

For a pattern named `<pattern>` (snake_case, e.g., `pubsub_with_topics`):

| File | Purpose |
|------|---------|
| `patterns/NN_<pattern>.go` | Implementation. `package main` with `func main()` demoing usage. **First line must be `//go:build ignore`**. |
| `patterns/NN_<pattern>.md` | When/why doc, rendered above the editor. |

`NN` is the next available numeric prefix (currently 01–21).

## Pattern selection criteria

A good pattern entry:
- Has a name interviewers will recognize (errgroup, circuit breaker, builder, repository, etc.)
- Solves a recurring problem with a teachable shape
- Has at least one stdlib or widely-used library precedent (cite it in the doc)
- Doesn't overlap heavily with an existing pattern — see "Check for duplication" below

Anti-patterns:
- A specific library wrapper (that's not a pattern, that's API docs)
- One pattern dressed as two — if you'd describe the new one as "the X pattern but with Y", consider expanding the existing X instead

## Check for duplication

Before adding, scan `patterns/` for overlap. Past examples of overlaps that were caught:
- Map/Filter/Reduce vs. basics/16_generics (the basics file was trimmed)
- Manual fan-out vs. patterns/16_fan_in_fan_out
- Sentinel/typed errors vs. basics/10_errors (basics now points to the pattern)

If the new pattern would overlap meaningfully, prefer one of: (a) extending the existing pattern's `.md`, (b) cross-referencing both files, (c) explicitly contrasting them in the new pattern's "When to use" section.

## .go file template

```go
//go:build ignore

package main

import (
    "fmt"
)

// PATTERN NAME
// ============
// 2-4 line description of what problem this solves.
//
// When to use this pattern (call out the canonical case).
// When NOT to use it (and what to use instead).

// implementation: keep types and methods small enough to fit on a screen.

func main() {
    // demonstrate the pattern with a tight example that prints output.
}
```

Style rules:
- Inline doc comments at the top should mirror the `.md` "When to use" / "When NOT" sections at a glance.
- The `main()` should be a self-contained demo: build the thing, exercise it, print results.
- Prefer one focused example over three half-examples. Quality > quantity.

## .md file template

```markdown
# <Pattern Name>

<1-2 sentences: what problem it solves.>

## The problem

<2-4 lines articulating the pain this pattern eliminates. Optional but
recommended for non-obvious patterns.>

## The shape

​```go
// the minimal code sketch — not the full implementation, just the
// signature/structure a reader would memorize for a whiteboard.
​```

## When to use

- <case>
- <case>

## When NOT to use

- <case>
- <case>

## Real-world

<1-3 places this pattern appears: stdlib, popular libraries, the user's own infra.>

## Interview frame

<The interview question this pattern answers. e.g., "Design a rate limiter" → token bucket.>
```

Sections to add only when relevant:
- `## Common pitfall` — if there's a specific footgun
- `## Variants` — if there are 2-3 distinct shapes worth comparing
- `## Pattern vs. <other pattern>` — when adjacency could confuse readers

## After creating files

1. Verify: `go run patterns/NN_<pattern>.go` produces clean output.
2. Verify whole-module build: `go build ./...`
3. The web app auto-discovers it.

## Reference

See `patterns/01_functional_options.{go,md}` for the canonical exemplar and `patterns/13_errgroup.{go,md}` for one with a "production version" framing. CLAUDE.md has broader conventions.
