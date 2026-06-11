# Fan-Out / Fan-In

Lower-level channel-composition pattern that worker pools and pipelines are built on.

## The two halves

- **Fan-out**: one input channel → N concurrent processors, each with its own output channel.
- **Fan-in**: N output channels → one merged channel.

## When to use directly

- Heterogeneous parallelism — different processors doing different work, all feeding one consumer.
- Composing existing pipelines that already have channels.
- Pipelines where you want N copies of an expensive stage in the middle.

## When a worker pool is better

- Homogeneous work items, all processed the same way → worker pool (`patterns/02`) is more direct.

## Generic shape

```go
func fanOut[T, R any](in <-chan T, n int, fn func(T) R) []<-chan R
func fanIn[R any](chans ...<-chan R) <-chan R
```

These two primitives compose to build any concurrent pipeline you'd actually want.

## Worth knowing

- The fan-in merger must close its output when all inputs close — that requires waiting on all input channels (`sync.WaitGroup`).
- Output order from fan-out is **not** preserved — items are processed in parallel and merged in completion order. Add per-item indices if you need to reassemble.
