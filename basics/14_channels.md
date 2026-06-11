# Channels

Typed pipes for goroutine communication. "Don't communicate by sharing memory; share memory by communicating."

## Worth knowing

- `make(chan T)` is unbuffered — sends block until a receiver is ready (and vice versa).
- `make(chan T, n)` is buffered — sends don't block until the buffer is full.
- Closing a channel: `close(ch)`. Receivers see the zero value with `ok=false`: `v, ok := <-ch`.
- Sending on a closed channel **panics**. Closing twice panics. Only the sender should close.
- `range ch` reads until the channel closes.

## Select

Multiplex on multiple channel operations:

```go
select {
case v := <-ch1: ...
case ch2 <- x:   ...
case <-ctx.Done(): return
default:         ...  // non-blocking
}
```

`default` makes the select non-blocking. Without it, select blocks until one case fires.

## Common gotchas

- Receiving from a nil channel blocks **forever**. Sometimes intentional (disable a case in select); usually a bug.
- Sending to a nil channel also blocks forever.
- Buffered channels can hide synchronization bugs — start unbuffered, add buffering only with a reason.

## Patterns

- **Done channel**: `done := make(chan struct{})` + `close(done)` broadcasts to all receivers.
- **Fan-in**: merge multiple channels into one (`patterns/16_fan_in_fan_out.go`).
- **Pipeline**: chain stages, each owning its output channel (`patterns/10_iterator_and_pipeline.go`).
