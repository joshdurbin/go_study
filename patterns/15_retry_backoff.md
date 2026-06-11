# Retry with Exponential Backoff + Jitter

Don't kill a transient failure with an immediate retry. Wait, increase the wait, and randomize.

## The recipe

1. **Exponential backoff** — wait `base * 2^attempt` after attempt N.
2. **Cap the backoff** — don't grow unbounded.
3. **Full jitter** — pick uniformly in `[0, window)` so concurrent retries don't synchronize.
4. **Respect context cancellation** — never sleep past `ctx.Done()`.
5. **Distinguish retryable from terminal errors** — don't retry on 4xx, do retry on 5xx.

## Why jitter

If 1000 clients fail at the same instant and all back off the same duration, they retry **simultaneously** — the dependency sees the same load that caused the original failure. Random jitter breaks the lockstep.

## When to use

- Network calls (DNS, HTTP, gRPC).
- Database/cache reconnects.
- Anywhere "this might just be a transient blip" applies.

## When NOT to use

- Idempotent failures with deterministic causes (validation errors, malformed input). Retrying won't help.
- Inside a hot loop with no real backoff — that's a busy-wait dressed up.

## Worth knowing

- AWS docs popularized the "full jitter" recipe. There are variants (decorrelated jitter, equal jitter) — full jitter is the default unless you have a reason.
- Pair with a circuit breaker (`patterns/07`) when the dependency stays down for long.
