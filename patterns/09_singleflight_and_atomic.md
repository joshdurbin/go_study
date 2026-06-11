# Singleflight + Atomic

Two distinct patterns for handling concurrent access efficiently.

## Singleflight

**Problem:** 100 goroutines all call `fetchUser(42)` simultaneously (cache miss).
**Without:** 100 identical DB queries fire.
**With:** 1 query fires; all 100 callers get the same result.

```go
result, err, _ := group.Do(key, func() (any, error) {
    return expensiveFetch(key)
})
```

`golang.org/x/sync/singleflight` is the standard implementation.

## When singleflight wins

- Cache stampede protection (key just expired, all requests miss simultaneously).
- Expensive idempotent operations (config fetch, topology lookup).
- Thundering herd on cold start.

## Atomic

**Problem:** lock contention on simple counters / flags.
**Solution:** `sync/atomic` (and `atomic.Int64` etc. since 1.19) — lock-free via CAS.

```go
var requests atomic.Int64
requests.Add(1)
n := requests.Load()
```

## When atomic wins

- High-frequency counters (request counts, hit/miss tracking).
- Feature flags toggled rarely, read constantly.
- Lock-free state machines (compare-and-swap state transitions).

## When NOT to use atomic

- Anything involving multiple variables that need to update together — that's a mutex.
- Anything more complex than read/write/CAS one word — the moment you reach for `atomic.Value` with complex types, prefer a mutex for clarity.
