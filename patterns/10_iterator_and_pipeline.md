# Iterator and Pipeline

Two related patterns for streaming data.

## Iterator

Abstract sequential access without exposing the underlying collection.

- **Callback** (`func(yield func(T) bool)`) — pre-Go-1.22 norm, and the shape Go 1.22's `iter` package formalized.
- **Channel** — caller ranges over a channel; simple but each iterator costs a goroutine.
- **Stateful struct** — explicit `Next()` / `Value()` (`sql.Rows`, `bufio.Scanner`).

## Pipeline

Chain stages where each stage reads from one channel and writes to another:

```go
nums := generate(1, 2, 3)        // chan int
squared := square(nums)          // chan int
for v := range squared { ... }
```

Each stage owns its output channel and closes it when done. The receiver knows the pipeline is finished when its input channel closes.

## When to use

- **Iterator**: anywhere you have a "stream of T" that callers want to consume lazily.
- **Pipeline**: multi-stage transformation where each stage can run concurrently.

## When NOT to use a pipeline

- You're processing a small in-memory slice — just chain function calls. Pipelines are for genuinely streaming data.

## Go 1.22 iter

Range-over-func makes iterators feel native: `for v := range collection.All { ... }`. New stdlib types are adopting it — worth recognizing.
