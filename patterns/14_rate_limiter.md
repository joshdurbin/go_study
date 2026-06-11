# Rate Limiter (Token Bucket)

Enforce "at most R requests per second, with burst capacity B."

## Token bucket vs. leaky bucket

- **Token bucket**: bucket holds B tokens; refills at R/sec; each request consumes one. Allows short bursts up to B then enforces R long-term. Most APIs use this.
- **Leaky bucket**: enforces a strict, even rate — no bursts. Useful when downstream can't absorb spikes at all.

## Implementation trick

Don't run a refill goroutine. Compute elapsed time on each `Allow()` call and add proportional tokens.

```go
elapsed := time.Since(lastRefill).Seconds()
tokens = min(capacity, tokens + elapsed * refillRate)
lastRefill = now
if tokens >= 1 { tokens--; return true }
return false
```

## When to use

- Client-side: throttling outbound calls to a rate-limited API.
- Server-side: per-user/per-IP rate limiting.
- Spike protection in front of expensive operations.

## When NOT to use

- Strict per-second quotas that must reset on a wall-clock boundary — token bucket is rolling, not aligned.
- Distributed rate limiting across many processes — needs a shared store (Redis, etc.). The in-process bucket isn't enough.

## Real-world

`golang.org/x/time/rate` is the stdlib-quality token bucket. Use it directly unless you need custom semantics.

## Interview frame

"Design a rate limiter" → token bucket is the right answer. Bonus points for naming both implementations and explaining the trade-off.
