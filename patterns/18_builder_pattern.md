# Builder Pattern

Accumulate state across chained method calls, then validate and emit the final object via `Build()`.

## Builder vs. Functional Options

Both solve "constructor with many configurable parts." Pick by:

- **Functional options** (`patterns/01`): fields are independent, no ordering constraints, validation per-option. The default for Go APIs.
- **Builder**: construction has steps that interact, cross-field validation at the end, or you want stateful chaining (e.g., a query DSL where `Where` calls accumulate).

## The shape

```go
q, err := From("users").
    Where("active = true").
    Where("age > 18").
    Limit(50).
    Build()
```

Each step returns `*QueryBuilder` so calls chain. Errors captured along the way surface at `Build()`.

## When to use

- DSLs and query builders.
- Multi-step construction with validation that needs the whole config.
- Configuration where order matters or steps interact.

## When NOT to use

- Independent options — functional options is cleaner and more idiomatic.
- You're tempted to expose the builder publicly for a type with two settings. That's a config struct, not a builder.

## Worth knowing

The "first error wins" pattern (record error in the builder, return it at Build) keeps the chain readable. Without it, every method returns `(*Builder, error)` and chaining dies.
