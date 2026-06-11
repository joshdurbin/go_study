# Unbuffered vs buffered channels

## The question

When should I use `make(chan T)` vs `make(chan T, N)`?

## The answer

- **Unbuffered** (`make(chan T)`): send and receive synchronize. The sender blocks until a receiver is ready, and vice versa. It's a rendezvous — both goroutines reach the channel before either proceeds.
- **Buffered** (`make(chan T, N)`): the channel holds up to N items. Sends only block when the buffer is full; receives only block when it's empty. Producer and consumer are decoupled in time as long as the buffer absorbs the burst.
- Pick unbuffered when you want **synchronization** (handoff, signal, "I'm done").
- Pick buffered when you want **decoupling** (work queue, rate smoothing, fan-out).
- Default to unbuffered. Add buffering only when profiling shows backpressure or you can prove the buffer is bounded by something real (e.g., number of pending HTTP requests).

## The gotcha

A buffered channel hides bugs. With `make(chan T, 1000)` a slow consumer can keep up "for now," masking the real problem until production traffic doubles. An unbuffered channel surfaces the contention immediately. Also: `make(chan T, 1)` is a different beast — it's a one-slot mailbox often used for "latest value wins" or non-blocking signals.

## In code

See the runnable demo in this file. Key output:
- `unbuffered send took ~30ms (waited for receiver)`
- `buffered (cap=3) 3 sends took ~µs (no receiver yet)`
- `buffered send when full took ~20ms`

## Related

- [[goroutine-leak-patterns]] — sends/receives that never complete
- [[select-statement]] — multiplexing channel ops
- [[fan-in-fan-out]] — design pattern using buffered channels
