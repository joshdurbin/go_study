# Interfaces

A set of method signatures. Types satisfy interfaces **implicitly** — there's no `implements` keyword. If a type has the methods, it satisfies the interface.

## Worth knowing

- Small interfaces are idiomatic. `io.Reader` has one method. The standard wisdom: "The bigger the interface, the weaker the abstraction."
- The empty interface `interface{}` (or its alias `any` since Go 1.18) accepts any value.
- Interface values are two words internally: (type, value). A nil interface is `(nil, nil)`. An interface holding a nil pointer is `(*T, nil)` — **not** nil.
- Type assertions: `v, ok := x.(MyType)` — never use the no-ok form unless you're sure, or you'll panic on mismatch.

## When to define one

When more than one type will satisfy it, OR when you need to mock it for tests. Don't define interfaces speculatively — that's a Java reflex Go doesn't share.

## The nil-interface gotcha

```go
var p *MyError = nil
var err error = p  // err is NOT nil — it holds (*MyError, nil)
```

This burns everyone once. Return `nil` directly, not a typed nil pointer.
