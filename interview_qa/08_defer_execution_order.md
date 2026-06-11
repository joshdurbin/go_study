# Defer execution order

## The question

What order do multiple `defer`s run in? When are their arguments evaluated? Can a deferred function change the return value?

## The answer

- Defers run in **LIFO** order — last deferred, first to execute. Think of a stack.
- **Arguments are evaluated when the defer statement executes**, not when it runs. `defer fmt.Println(x)` snapshots `x` at defer time. A `defer func() { fmt.Println(x) }()` captures `x` as a closure variable and sees the latest value.
- A deferred function CAN modify **named return values**. This is how `defer recover()` patterns work and how you can post-process return values.
- `defer` in a loop accumulates — every iteration adds another deferred call. They all fire when the enclosing function returns, not at end-of-iteration. For per-iteration cleanup, refactor the body into a helper function.
- Defers fire even on panic — that's the whole point of `recover`.

## The gotcha

`defer f.Close()` inside a `for` over many files holds every file open until the function returns. In a long-lived function this leaks descriptors. Also: people forget the argument-evaluation rule and assume `defer log.Printf("status=%d", status)` will log the final status. It logs whatever `status` was at defer-time.

## In code

See the runnable demo in this file. Key output:
- LIFO: `3 2 1`
- args captured: `current x: 20` then `deferred x: 10`
- closure: `closure x: 20`
- named return mutation: `wrap() = wrapped(inner)`
- loop: `loop body done` then `2 1 0`

## Related

- [[defer-panic-recover]] — the broader trio
- [[goroutine-leak-patterns]] — deferring `cancel()` is the standard idiom
- [[named-return-values]] — when to use them, when not to
