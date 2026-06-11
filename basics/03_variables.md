# Variables

Two declaration forms: `var x int = 5` (explicit) and `x := 5` (short, function-scope only with inference).

## Worth knowing

- `:=` requires at least one new variable on the left; reusing existing names is fine if at least one is new.
- Unused local variables are a **compile error** — Go forces you to delete dead code.
- Multiple return values fall out naturally: `a, b := f()`.
- `_` is the blank identifier — discard a value you don't need.

## Common gotcha

`:=` inside a new scope (loop, if-block) creates a NEW variable that shadows the outer one. Use `=` to assign to the outer.
