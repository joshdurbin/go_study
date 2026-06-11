# Circuit Breaker

Stop hammering a failing dependency. After N failures, the breaker "opens" and short-circuits subsequent calls; after a cool-down it goes "half-open" and tries one call to see if recovery happened.

## Three states

- **Closed**: normal operation. Failures increment a counter.
- **Open**: short-circuit — return an error immediately without calling the dependency.
- **Half-open**: a single probe call is allowed. Success → close. Failure → re-open.

## When to use

- Downstream service that's known to flap or have outages.
- Anywhere a thundering herd of retries would make a bad situation worse.
- In front of slow dependencies where you'd rather fail fast than queue up.

## When NOT to use

- Idempotent calls with cheap, quick retries — exponential backoff (`patterns/15`) may be enough.
- Internal in-process calls — circuit breakers are an inter-service tool.

## Real-world

`github.com/sony/gobreaker` is the standard Go implementation. Netflix's Hystrix popularized the pattern. Service meshes (Istio, Linkerd) provide it at the network layer.

## Interview frame

"How do you protect a service from a failing dependency?" → circuit breaker is the right vocab. Bonus: explain the half-open state — that's the part candidates often forget.
