# Memoization

Cache the result of an expensive deterministic function, keyed by its arguments.

## The shape

```go
type Memo[K comparable, V any] struct {
    mu    sync.RWMutex
    cache map[K]V
    fn    func(K) V
}

func (m *Memo[K, V]) Get(k K) V {
    // RLock fast path → Lock + double-check on miss → compute → store
}
```

Double-checked locking: RLock for the common case (hit), upgrade to Lock only on miss, re-check after upgrading.

## Memoization vs. caching vs. singleflight

- **Memoization**: permanent, in-process cache. Bridges directly to top-down DP.
- **LRU cache**: bounded with eviction. See `interview_study/.../lru_cache.go`.
- **Singleflight**: dedupe concurrent CALLS for the same key without keeping the result. See `patterns/09`.

These compose: a real production cache often layers singleflight (for stampede protection) over a TTL'd cache (for staleness control) over memoization (for raw hit-rate).

## When to use

- Expensive pure function called repeatedly with the same args.
- Top-down dynamic programming converted to a tabular solution.
- Repeated transformations in tight loops (parsed configs, compiled regexes).

## When NOT to use

- Non-deterministic function. The cache will lie.
- Argument space is huge and rarely repeats — you'll fill memory with one-time results.
- You need bounded memory — use an LRU instead.
