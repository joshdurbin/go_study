# Goroutine Patterns

Concrete intro examples for the four most common channel/sync coordinations.

## The four shapes shown here

1. **Fan-out** — distribute work across N workers reading from a shared channel.
2. **Pipeline** — chain stages where stage K's output channel is stage K+1's input.
3. **Semaphore** — bounded concurrency using a buffered channel as a token bucket.
4. **Errgroup-style waiter** — wait for N tasks, propagate the first error.

## Worth knowing

- Always know when each goroutine ends. The most common bug is a goroutine blocked forever on a channel.
- Always close channels from the SENDER side, never the receiver. Closing from the receiver is a panic.
- Buffered channels can mask synchronization bugs. Start unbuffered and add buffering only with a measured reason.

## Production versions

The examples in this file are deliberately concrete and minimal. For the generic, composable versions:

- `patterns/02_worker_pool.go` — worker pool with results channel
- `patterns/13_errgroup.go` — first-error + sibling cancellation
- `patterns/16_fan_in_fan_out.go` — generic fan-out + fan-in primitives
- `patterns/10_iterator_and_pipeline.go` — pipeline composition (incl. Go 1.22 iter)
