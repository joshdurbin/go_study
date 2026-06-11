# Nil interface vs nil pointer

## The question

Why does this print `false`? `var p *T = nil; var i interface{} = p; fmt.Println(i == nil)`.

## The answer

- An interface value is a two-word header: `(type, value)`.
- `i == nil` is true only when BOTH the type and value words are zero.
- Assigning a typed nil pointer fills the type word with `*T`, so the interface is non-nil even though the pointer it carries is nil.
- The fix: return literal `nil` (untyped), or compare the concrete pointer with `nil` before wrapping.

## The gotcha

Functions returning `error` that do `var e *MyError; ... return e` always return a non-nil error, even on the success path. Every `if err != nil` after that call fires. This bug ships to prod constantly.

## In code

See the runnable demo in this file. Key output:
- `p == nil: true`
- `i == nil: false`
- `ERROR returned (but the underlying pointer is nil!): ...`
- `err2 == nil: true`

## Related

- [[type-assertion-vs-type-switch]] — extracting the concrete value from an interface
- [[error-wrapping-unwrapping]] — `errors.Is` / `errors.As` interact with this
- [[interface-internals]] — itab and the (type, value) pair layout
