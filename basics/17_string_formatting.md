# String Formatting

`fmt` is the standard formatter — `Printf`, `Sprintf`, `Errorf`, and friends.

## Verbs you'll use constantly

- `%v` — default format. `%+v` adds field names for structs. `%#v` Go-syntax representation.
- `%d` — integer. `%f` — float. `%.2f` — 2 decimal places.
- `%s` — string. `%q` — quoted string (escapes special chars).
- `%t` — bool. `%p` — pointer address. `%T` — the type itself.
- `%w` — **wrap** an error (only valid in `fmt.Errorf`). Preserves the chain for `errors.Is` / `errors.As`.

## Width and padding

- `%5d` — right-align in width 5. `%-5d` — left-align. `%05d` — zero-pad.
- `%10.3f` — width 10, 3 decimals.

## Building strings efficiently

For lots of concatenation, use `strings.Builder` — `+` allocates a new string each time.

```go
var b strings.Builder
for _, s := range parts {
    b.WriteString(s)
}
return b.String()
```

## Worth knowing

- `fmt.Errorf("...: %w", err)` is the only way to wrap an error so `errors.Is` finds it.
- `fmt.Println(x)` calls `Error()` on errors and `String()` on stringers automatically.
