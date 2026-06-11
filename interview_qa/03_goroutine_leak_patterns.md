# Goroutine leak patterns

## The question

Show me how to leak a goroutine. Now show me how to fix it.

## The answer

- A goroutine leaks when it's blocked on a channel operation that will never complete (send to a full/never-read channel, receive from a never-sent/never-closed channel).
- The runtime can't garbage-collect a live goroutine; leaks accumulate until OOM.
- Three common patterns:
  - Consumer waits on `<-ch`; producer panics or returns early without closing.
  - Worker sends on `ch <- v` but nobody is reading anymore.
  - `time.After` in a tight loop creates timers that don't fire — use `time.NewTimer` + `Stop()`.
- Fixes:
  - `select` with `case <-ctx.Done():` — always have an escape hatch.
  - Owner closes the channel on shutdown so receivers' `range` loop ends.
  - Use `errgroup.Group` to coordinate; it propagates cancellation.

## The gotcha

The leak is silent. `go test` passes, the request returns 200, the binary runs for days. Only `pprof /debug/pprof/goroutine` or a steadily climbing `runtime.NumGoroutine()` reveals it. Make it a habit: every goroutine you spawn must have a documented way to exit.

## In code

See the runnable demo in this file. Key output:
- `goroutines after 5 leaks: N` — climbs by 5
- `goroutines after cancel:   N` — drops back after cancel propagates

## Related

- [[unbuffered-vs-buffered-channels]] — buffered channels hide some leaks, expose others
- [[context-propagation]] — `ctx.Done()` is the standard cancellation channel
- [[errgroup-pattern]] — coordinated goroutine lifecycle
