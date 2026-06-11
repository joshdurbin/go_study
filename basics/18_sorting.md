# Sorting

`sort` package + Go 1.21's `slices.Sort` cover almost every case.

## Built-ins

- `sort.Ints(s)`, `sort.Strings(s)`, `sort.Float64s(s)` — sort native types.
- `slices.Sort(s)` (Go 1.21+, generic) — same thing, prefer this for new code.
- `sort.Slice(s, less)` — custom comparator over a slice.
- `sort.Stable(...)` — preserves original order of equal elements (use when ties matter).

## Custom sort

```go
sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age
})
```

For complex sorts, define a type with `Len/Less/Swap` (the `sort.Interface`).

## Searching

`sort.Search(n, pred)` does binary search via a predicate — useful for "first index where condition is true".

## Worth knowing

- The default `sort` is **not stable**. Use `sort.SliceStable` if order of equals matters.
- For huge data, consider `sort.Slice` + a separate index slice to avoid copying large structs.
