# sync Patterns

The less-obvious parts of the `sync` package: `sync.Map`, `sync.Pool`, `sync.Cond`.

## sync.Map

Concurrent map with no external locking. **Not** a drop-in replacement for `map[K]V` + Mutex — it's optimized for very specific access patterns.

- Use when: **write once, read many** keys (cache, registration table).
- Don't use when: keys are frequently written/updated — a regular map + RWMutex is faster.

## sync.Pool

Per-CPU caches of reusable objects. Reduces GC pressure for short-lived objects.

- Use when: you allocate and discard the same shape of buffer/struct millions of times per second.
- Don't use when: the pool would hold large objects forever — sync.Pool can evict anything anytime, but pooled objects still keep memory alive between collections.
- Always reset pooled objects when you put them back (`buf.Reset()`).

## sync.Cond

Signal/broadcast over a condition. Rarely the right answer — channels usually do the job more clearly.

- Use when: you specifically need broadcast semantics that channels can't easily express (waking N waiters on a state change).
- Otherwise: a `chan struct{}` or `close()` broadcasts to all receivers and is simpler.

## Interview frame

If asked "when would you use sync.Pool?" — gc pressure reduction in hot paths. If asked about sync.Cond — you can probably do it with channels; here are the niche cases where you can't.
