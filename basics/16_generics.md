# Generics

Added in Go 1.18. Lets you write functions and types that work over multiple types without `interface{}` and reflection.

## Syntax

- Type parameters: `func Map[T, U any](s []T, fn func(T) U) []U`
- Constraints: `func Sum[T int | float64](s []T) T` — restricts T to listed types.
- Constraints can be interfaces: `type Number interface { ~int | ~float64 }`. The `~` means "or any type whose underlying type is this".
- Generic types: `type Stack[T any] struct { items []T }`.

## When to reach for them

- The same logic over multiple concrete types (numeric ops, collection helpers).
- Type-safe collections (a stack, queue, set).
- Library code where callers shouldn't need type assertions.

## When NOT to

- "Just in case" generics that have one caller — concrete code is clearer.
- Where an interface already works (e.g., `io.Reader`). Interfaces are the older, simpler tool.

## Worth knowing

The stdlib `slices` and `maps` packages (Go 1.21+) provide many generic helpers — `slices.Contains`, `slices.Index`, `maps.Keys`. Check there before writing your own.

See `patterns/19_generic_helpers.go` for Map/Filter/Reduce/GroupBy/FlatMap.
