# Defer / Panic / Recover

Three rarely-used (but interview-favorite) features.

## Defer

`defer fn()` schedules `fn` to run when the surrounding function returns — including via panic. Use it for resource cleanup: `defer file.Close()`, `defer mu.Unlock()`.

- Deferred calls run in **LIFO** order (last deferred, first to run).
- Arguments are evaluated AT defer time: `defer fmt.Println(x)` captures `x` then; `fn()` is what runs later.
- Defer in a loop accumulates — `for ... { defer f.Close() }` postpones every close until the function returns. Refactor or use an explicit close.

## Panic

`panic(value)` unwinds the stack, running deferred functions on the way up. If unrecovered, the program crashes with a stack trace.

- Reserve panic for **truly impossible** situations (invariant violated, programming bug).
- For expected failures, return an `error`.

## Recover

`recover()` returns the panic value if called inside a deferred function during unwinding; otherwise returns nil. It STOPS the unwind.

```go
defer func() {
    if r := recover(); r != nil {
        log.Println("recovered:", r)
    }
}()
```

Common production use: HTTP servers' top-level middleware catches panics per request so one bad handler doesn't crash the whole process.
