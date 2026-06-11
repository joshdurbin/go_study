# Errors

Errors are **values**, not exceptions. They flow through return values, not stack unwinding.

## Worth knowing

- `error` is just an interface with one method: `Error() string`.
- Three patterns work together:
  - **Sentinels**: `var ErrNotFound = errors.New("not found")` — compare with `errors.Is`.
  - **Custom types**: implement `Error()` to carry data — match with `errors.As`.
  - **Wrapping**: `fmt.Errorf("layer X: %w", err)` preserves the chain.
- `panic` is for **truly unexpected** situations (programming bugs, impossible states). Not error handling.
- `recover` only works inside a `defer`'d function.

## Style

- `if err != nil { return err }` after every fallible call. It's verbose but explicit.
- Don't log AND return — pick one. Logging at every layer creates noise; let the top of the call stack decide what to do.
- For library code, return typed/sentinel errors callers can match on. For app code, plain `fmt.Errorf` is usually fine.

See `patterns/20_error_wrapping.go` for the production-grade error model.
