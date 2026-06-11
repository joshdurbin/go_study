# iota & Enums

Go has no `enum` keyword. The idiom is a named integer type plus a `const` block using `iota` — a compile-time counter that resets each block and increments by 1 per line.

## Worth knowing

- `iota` starts at 0 and increments once per `ConstSpec` (one line) in the block. The expression on the first line is repeated implicitly on later lines.
- `1 << iota` gives powers of two — the idiom for bit flags you combine with `|` and test with `&`.
- Use `_` to skip a value (commonly to discard 0 when zero should mean "unset").
- Define the underlying type explicitly (`type Status int`) — gives type safety and a place to hang methods.

## The Stringer interface

```go
type Stringer interface { String() string }
```

Implement `String() string` on your enum type and `fmt` packages will use it automatically when formatting with `%s` or `%v`. Without it, you'd see raw integers in logs.

## go:generate stringer

Writing `String()` by hand is tedious and gets stale. Add `//go:generate stringer -type=Status` above the const block, then `go generate ./...` emits a `status_string.go` with a correct, exhaustive switch. Install once with `go install golang.org/x/tools/cmd/stringer@latest`.

## Common gotcha

The zero value of any int type is 0. If your first enum constant is at iota=0, the zero value of a freshly declared variable IS that constant — which may not be what you want. For state machines where "unset" should be distinct from "first state," start the enum at 1 (`_ = iota` to discard 0).

## Interview frame

If asked about modeling state: mention typed iota constants, Stringer for logging, and `go:generate stringer` for upkeep. If asked about flags/permissions: `1 << iota` plus bitwise ops.
