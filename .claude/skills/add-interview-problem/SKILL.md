---
name: add-interview-problem
description: Add a new algorithm interview problem to the Go study guide under interview_study/ — creates the full four-file set (problem statement, solution, progressive hints, editor starter) and places it in the right block/tier. Use when the user wants to add a new interview practice problem.
---

# Add an interview problem

Creates a full interview lesson with four files. Tier and block placement matter — the skill picks them based on the problem's character.

## Files to create

For a problem named `<problem>` (snake_case, e.g., `longest_palindrome_substring`) in directory `<dir>`:

| File | Purpose |
|------|---------|
| `<dir>/<problem>.go`         | Reference solution. `package main`, `func main()` with example calls and expected output as comments. **First line: `//go:build ignore`**. |
| `<dir>/<problem>.md`         | Problem statement (markdown). What to do, constraints, examples. |
| `<dir>/<problem>.steps.md`   | Progressive hints. Tier determines count (see below). |
| `<dir>/<problem>.starter.go` | Editor scaffold: function signatures with `// TODO: implement` bodies + `main()` with test cases. **First line: `//go:build ignore`**. |

## Choose the block and tier

The block reflects the **technique** the problem teaches:

| Block | Directory | Use for |
|-------|-----------|---------|
| Warmups (Easy) | `interview_study/block0_warmups/1_easy/` | Single-loop / single-pass problems, basic data-structure use |
| Warmups (Medium) | `interview_study/block0_warmups/2_medium/` | Classic medium LeetCode-style — sliding window, stack, multi-pass |
| Warmups (Hard) | `interview_study/block0_warmups/3_hard/` | Multi-data-structure (LRU), tricky binary search, hard DP |
| Go Fundamentals | `interview_study/block1_go_fundamentals/` | Tests Go-specific knowledge (heap.Interface, map idioms, slice mechanics) |
| Two Pointers / Window | `interview_study/block2_two_pointers_sliding_window/` | Two-pointer or sliding-window technique is the key |
| Hashmaps / Frequency | `interview_study/block3_hashmaps_frequency/` | Frequency counts, prefix sums + map, grouping |
| Binary Search | `interview_study/block4_binary_search/` | Search over indices or "search the answer" |
| Trees / Graphs | `interview_study/block5_trees_graphs/` | DFS/BFS, traversal, LCA, topo sort, Dijkstra |
| Dynamic Programming | `interview_study/block6_dynamic_programming/` | 1-D / 2-D DP, knapsack, edit distance |

The TIER for warmups is independent of block — for blocks 1–6, treat as medium by default, except block6 (DP) which gets hard treatment in hint count.

## Hint count by tier

| Tier              | Hints | Shape |
|-------------------|-------|-------|
| Easy              | 2     | nudge → approach |
| Medium            | 3     | nudge → approach → refinement/pitfall |
| Hard              | 4     | nudge → key insight → recurrence/structure → implementation detail |

Each hint is ~1-3 sentences of prose + a 3-10 line fenced Go code block that builds incrementally. The last hint should reveal ~80% of the solution but stop short of identical code.

## `.go` (solution) template

```go
//go:build ignore

package main

import "fmt"

// <funcName> <one-line spec — what it returns and any constraints>.
// O(...) time, O(...) space.
func funcName(args ...) returnType {
    // clean, idiomatic Go. Comments only where logic is non-obvious.
}

func main() {
    fmt.Println(funcName(...)) // expected output as comment
    fmt.Println(funcName(...)) // another case
}
```

## `.md` (problem statement) template

```markdown
PROBLEM: <Title>
================
<2-4 line problem description. State inputs, outputs, constraints.>

Why it's <tier>: <1-2 sentences on what the problem teaches.>

Example:
  Input:  ...
  Output: ...
```

The format is loose — existing problems use setext headers (`===`) and indented code blocks. Both render as markdown. New problems can use `## Header` and ` ```... ``` ` fences if you prefer; both render fine.

## `.steps.md` (hints) template

```markdown
## Hint 1
<Nudge: name the pattern, point at the data structure choice.>

​```go
// minimal code skeleton — variable declarations and loop shape
​```

## Hint 2
<Approach: what the body of the loop / function does.>

​```go
// inner logic that fills in the skeleton above
​```

## Hint 3
<Refinement: edge cases, off-by-one, the pitfall to avoid.>

​```go
// the corrected/final piece
​```

## Hint 4 (hard only)
<Implementation detail: stale-entry skip in Dijkstra, sentinel choice in DP, etc.>

​```go
// the last piece a reader needs
​```
```

## `.starter.go` template

```go
//go:build ignore

package main

import "fmt"

// <funcName> <same one-line spec as the solution — keep them in sync>.
// <Hint: name the technique to use, e.g., "Target O(n) time with sliding window.">
func funcName(args ...) returnType {
    // TODO: implement
    return /* zero value */
}

func main() {
    fmt.Println(funcName(...)) // expect <output>
    fmt.Println(funcName(...)) // expect <output>
}
```

For problems requiring helper types (LinkedList, TreeNode, custom data structure), keep the type definitions in the starter so the user doesn't have to redefine them. They're scaffolding, not part of the puzzle.

## Style rules across all four files

- **No emojis** anywhere — neither in code nor docs.
- **Tight narrative**: a learner is reading these under cognitive load. One sentence per hint sub-point. No filler.
- **Solution + starter agree on function signatures**. The user's edited starter should be drop-in compatible with the solution's `main()`.
- **Examples in `main()` match between starter and solution**. Same inputs, same expected outputs.
- **No dependencies outside the stdlib**. Anything more interesting belongs in `patterns/`.

## After creating files

1. Verify the solution compiles and runs: `go run <dir>/<problem>.go` — output should match the comments.
2. Verify the starter compiles (will print empty/nil results, that's expected): `go run <dir>/<problem>.starter.go`.
3. Verify whole-module build: `go build ./...`.
4. The web app picks it up on page refresh — no restart.

## Reference

Exemplars to mimic:
- Easy: `interview_study/block0_warmups/1_easy/two_sum.{go,md,steps.md,starter.go}`
- Medium: `interview_study/block0_warmups/2_medium/valid_parentheses.{...}`
- Hard: `interview_study/block0_warmups/3_hard/lru_cache.{...}`
- DP (4 hints): `interview_study/block6_dynamic_programming/edit_distance.{...}`

CLAUDE.md has the broader project conventions and file naming rules.
