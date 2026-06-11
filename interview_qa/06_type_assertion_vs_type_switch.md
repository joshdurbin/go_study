# Type assertion vs type switch

## The question

When do I use `v, ok := i.(T)` and when do I use `switch v := i.(type)`?

## The answer

- **Type assertion** (`i.(T)`) asks: is the interface value `i` of concrete type `T`? Two forms:
  - One-value: `v := i.(T)` — panics if `i` is not a `T`.
  - Comma-ok: `v, ok := i.(T)` — `ok` is false and `v` is the zero value on mismatch. Always prefer this unless you can prove the type.
- **Type switch** (`switch v := i.(type)`) dispatches over many possible types in one statement, binding `v` to the matched type inside each case. The `default` case still sees `v` typed as the original interface.
- Use assertion when you expect one specific type (often after a successful generic interface return).
- Use type switch when handling several types — printer functions, codec dispatch, AST walks.

## The gotcha

`i.(T)` without comma-ok panics with `interface conversion: ... is X, not T`. Library code should almost always use comma-ok. Also: a type switch with `case nil` actually catches a nil interface — useful for distinguishing "no value" from "zero value of some type." And `case *Cat, *Dog:` (multiple types per case) makes `v` keep the interface type — you can't access `Name` without further asserting.

## In code

See the runnable demo in this file. Key output:
- `string? true hello`
- `int?    false 0`
- `int: 42`, `string: go`, `cat: Mittens`, `dog: Rex`, `unknown: float64`, `nil`
- `recovered from panic: interface conversion: interface {} is string, not int`

## Related

- [[nil-interface-vs-nil-pointer]] — the (type, value) header underlying both
- [[error-wrapping-unwrapping]] — `errors.As` is a type assertion in disguise
- [[generics-vs-interfaces]] — when to reach for `any` + assertion vs `[T any]`
