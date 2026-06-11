# Functions

First-class values: pass them, return them, store them. Multiple return values are idiomatic, not an afterthought.

## Worth knowing

- Multiple returns enable the `(value, error)` convention used throughout stdlib.
- **Named return values** allow `return` with no arguments (a "naked return"). Useful with `defer` for error-modifying patterns, but don't overuse — it can hurt readability.
- Variadic: `func f(args ...int)` — `args` is a `[]int` inside.
- Closures capture variables by reference. Be careful around loop variables (fixed in Go 1.22+).
- Functions can be methods (have a receiver) or plain functions. Methods are functions with a `(receiver Type)` between `func` and the name.
