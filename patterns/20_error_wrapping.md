# Error Wrapping

Three idioms working together: sentinels, typed errors, and `%w` wrapping.

## Sentinel errors

```go
var ErrNotFound = errors.New("not found")
```

Compare via `errors.Is(err, ErrNotFound)` (it walks the wrap chain).

## Typed errors

```go
type ValidationError struct { Field, Msg string }
func (e *ValidationError) Error() string { ... }
```

Match via `errors.As(err, &ve)` — walks the chain looking for a typed value, binds it.

## Wrapping with %w

```go
return fmt.Errorf("getUser %q: %w", id, err)
```

`%w` preserves the chain — only valid in `fmt.Errorf`. `%s` or `%v` produce a NEW error with no chain.

## Anti-patterns

- **String matching on `err.Error()`**: fragile, breaks on rewording. Use `errors.Is` or `errors.As`.
- **Wrapping just to add context that's already in the stack**: noise. Wrap when you're crossing a layer boundary or adding genuinely new info.
- **Returning a typed nil pointer as `error`**: see the nil-interface gotcha in `basics/09_interfaces.md`.

## When to define which

- **Sentinel**: stable identity, no data (`io.EOF`, `sql.ErrNoRows`).
- **Typed**: carries data the caller needs (`ValidationError.Field`).
- **Wrapped plain**: adding context but the inner error already has identity.

## Worth knowing

Many stdlib errors are sentinels (`io.EOF`, `os.ErrNotExist`, `sql.ErrNoRows`) — `errors.Is` lets you check for them through any wrap chain. This composes cleanly with your own sentinel/typed errors.
