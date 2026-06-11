# For Loops

`for` is Go's only loop keyword. No `while`, no `do-while`.

## Forms

- Classic C-style: `for i := 0; i < n; i++ { ... }`
- While-style: `for cond { ... }`
- Infinite: `for { ... }` (exit with `break` or `return`)
- Range: `for i, v := range slice { ... }` — also works on maps, strings, channels, integers (Go 1.22+)

## Worth knowing

- Range over a string yields **rune indices and rune values**, not bytes. Bytes need `for i := 0; i < len(s); i++`.
- Range over a map has randomized order — never depend on it.
- Go 1.22+: `for i := range 10` ranges over `0..9`. Cleaner than the C form for fixed counts.
- Each iteration of a Go 1.22+ for loop creates a NEW loop variable — the pre-1.22 closure capture bug is gone.
