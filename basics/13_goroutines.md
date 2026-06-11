# Goroutines

`go fn()` launches a goroutine — a lightweight thread the Go runtime multiplexes onto OS threads. Spawning thousands is normal.

## Worth knowing

- A goroutine starts running as soon as it's scheduled. The launching code doesn't wait for it.
- When `main` returns, ALL goroutines die regardless of state. Use `sync.WaitGroup` or channels to coordinate.
- Goroutines are cheap (~2KB stack initially, grows on demand) — but not free. Don't launch one per tiny task; use a worker pool.

## Sync primitives

- `sync.WaitGroup`: wait for N goroutines to finish.
- `sync.Mutex` / `sync.RWMutex`: protect shared state.
- `sync.Once`: run something exactly once across goroutines.
- `sync.Cond`: rare — use channels instead in 95% of cases.

## Idioms

- **Always know how a goroutine ends.** Leaks come from goroutines blocked on a channel that's never sent to.
- Pass context through goroutines for cancellation. Don't rely on shared bool flags.
- `defer wg.Done()` at the top of a goroutine guarantees the counter decrements even on panic.

See `patterns/13_errgroup.go` and `patterns/16_fan_in_fan_out.go` for production patterns.
