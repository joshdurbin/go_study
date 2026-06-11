# Observer / Event Bus

One source, many listeners. In Go, idiomatically channels or callbacks — **not** the classic OOP Subject/Observer class hierarchy.

## Two implementations

- **Callback-based**: synchronous, simple, good for in-process events.
- **Channel-based**: asynchronous, supports backpressure, good for fan-out to goroutines.

## When to use

- A producer needs to notify N consumers that may grow/shrink over time.
- Plugin architectures where subscribers register at runtime.
- Cross-cutting "something interesting happened" notifications.

## When NOT to use

- Two specific consumers known at compile time — just call them.
- Persistence required across restarts — use a real message broker (NATS, Kafka).
- Strong ordering or exactly-once delivery semantics — also use a broker.

## Real-world

Prometheus metric registrations, Kubernetes informer event handlers, AlertManager subscribers — all variants of the pattern.

## Trade-off

The channel-based version has built-in backpressure (slow consumer blocks the producer) which is sometimes what you want and sometimes catastrophic. Buffered channels + non-blocking sends (with overflow logging) is the production answer.
