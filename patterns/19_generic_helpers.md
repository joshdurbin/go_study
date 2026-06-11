# Generic Map / Filter / Reduce

The functional-helper toolkit Go 1.18+ enables.

## Worth knowing

- The stdlib `slices` and `maps` packages (Go 1.21+) cover many cases: `slices.Contains`, `slices.Index`, `slices.SortFunc`, `maps.Keys`, etc. **Check there first** before writing your own.
- Generic Reduce shines because of its TWO type parameters: the accumulator type can differ from the element type, so you can reduce `[]User` to `int`, `map[string]int`, or anything else.
- `GroupBy` and `FlatMap` aren't in the stdlib yet — generics make them one-liners.

## Style guideline

In Go, an explicit for-loop is often clearer than chained Map/Filter/Reduce. Use these helpers when they **genuinely simplify**, not because they exist.

```go
// for-loop is clearer here
sum := 0
for _, x := range nums { sum += x }

// helper is clearer when chained
emails := Map(Filter(users, isActive), func(u User) string { return u.Email })
```

## Interview frame

"Show me a function that's only possible with generics" → Reduce with different accumulator type. Or a generic Stack/Queue. These demonstrate you've internalized generics beyond the obvious slice-of-T cases.

## When NOT to use

- One caller and one specific type — concrete code is friendlier.
- The closure passed to Map/Filter is non-trivial — extract a named function instead of writing complex inline lambdas.
