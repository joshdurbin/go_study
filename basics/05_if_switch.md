# If / Switch

Two distinguishing features vs. C-family languages: `if` can have an init clause, and `switch` doesn't fall through by default.

## If

```go
if v, err := fn(); err != nil {
    return err
}
// v is out of scope here
```

The init clause scopes the variable to the if/else block — keeps your variable lifetimes tight.

## Switch

- No `break` needed — cases don't fall through.
- Use `fallthrough` explicitly if you actually want C-style behavior.
- `switch { case cond1: ...; case cond2: ... }` (no operand) is the cleanest if/else-if chain.
- Type switch: `switch v := x.(type) { case int: ...; case string: ... }`.

## Worth knowing

Go style: prefer early returns and the init clause over deep nesting. `if err != nil { return err }` is a one-liner.
