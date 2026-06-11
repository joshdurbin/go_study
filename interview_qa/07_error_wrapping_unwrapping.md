# Error wrapping / unwrapping

## The question

What's the difference between `%v` and `%w` in `fmt.Errorf`? How do `errors.Is` and `errors.As` work?

## The answer

- `fmt.Errorf("...%w", err)` returns an error whose `Unwrap() error` method returns `err`. That's what "wrapping" is — a single-linked list of errors.
- `errors.Is(err, target)` walks `err.Unwrap()` repeatedly, comparing with `==` (or the target's `Is` method) at each step. Use it with **sentinel errors** (`var ErrNotFound = errors.New(...)`).
- `errors.As(err, &target)` walks the chain and assigns the first error of matching type into `target`. Use it with **typed errors** (`*ValidationError`).
- `%v` formats the error's text but does NOT preserve the chain — `errors.Is` won't find it through `%v`.
- Wrap when crossing layers, with enough context to debug. Don't double-wrap with no new info.

## The gotcha

`errors.Is` requires you to wrap with `%w`, not `%v`. Forget the `w` and every caller's `errors.Is(err, ErrNotFound)` silently returns false. Also: `errors.As` takes a **pointer** to a destination — `errors.As(err, &myErr)`, not `errors.As(err, myErr)`. The latter compiles only if `myErr` is itself a pointer and panics otherwise.

## In code

See the runnable demo in this file. Key output:
- `err: handleRequest: loadUser("u-42"): not found`
- `Is ErrNotFound?  true`
- `validation error on: email`
- `is os.ErrNotExist?  true`
- `Is ErrNotFound via %v?  false`

## Related

- [[nil-interface-vs-nil-pointer]] — `errors.As(nil-typed-pointer)` is a common bug
- [[sentinel-vs-typed-errors]] — when to use each style
- [[panic-vs-error]] — wrapping is for errors, not panics
