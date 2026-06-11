# Values and Types

Go is statically typed with explicit primitive widths (`int32`, `float64`, etc.) and a default `int` that's either 32 or 64 bits depending on the platform.

## Worth knowing

- Every type has a **zero value**: `0` for numerics, `""` for strings, `false` for bools, `nil` for pointers/slices/maps/channels/functions/interfaces.
- Untyped constants (`const x = 42`) take on whatever numeric type they need in context — `int`, `int64`, `float64` — without explicit conversion.
- Conversions between numeric types are **always explicit**: `float64(myInt)`. Go never silently widens.
- `string` is an immutable byte sequence (UTF-8 by convention); `len(s)` returns bytes, not runes.

## Common gotcha

`int` and `int64` are distinct types even when both are 64-bit. `int(x)` is required.
