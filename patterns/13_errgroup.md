# errgroup

Run N goroutines, wait for them all, return the first error, and cancel siblings when it happens. `sync.WaitGroup` alone gives you Wait but not the error propagation or cancellation.

## When to use

- Parallel fetches that should all succeed.
- Fan-out aggregation where any failure means abort.
- Anywhere "all of these must complete or we're done" applies.

## The shape

```go
g, ctx := errgroup.WithContext(parent)
for _, url := range urls {
    url := url
    g.Go(func() error {
        return fetch(ctx, url)
    })
}
if err := g.Wait(); err != nil { /* first error */ }
```

The shared `ctx` cancels when any goroutine errors — others see `ctx.Done()` and bail.

## When NOT to use

- "Best effort" parallelism where partial failures are OK — collect errors into a slice instead.
- Heterogeneous tasks where you need per-task status — use channels.

## Worth knowing

- `g.SetLimit(N)` bounds concurrency to N (Go 1.20+).
- The `errgroup` API is in `golang.org/x/sync/errgroup`. The pattern is ~30 lines — implementable from memory on a whiteboard.

## Production canon

`errgroup` is the answer to almost every "N parallel things, all must succeed" interview question.
