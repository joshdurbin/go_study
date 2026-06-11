# Slice aliasing footguns

## The question

What does `b := a[1:3]; b[0] = 99` do to `a`? When is `append` safe to use on a sub-slice without affecting the parent?

## The answer

- A slice is `(ptr, len, cap)` over a backing array. Slicing creates a new header but the same `ptr`.
- Mutations through any header are visible to every header pointing into the same region. `b[0] = 99` writes `a[1]`.
- `append` is sometimes safe and sometimes not: if `len(s) < cap(s)`, append writes in place into the shared array. If it would exceed cap, Go allocates a new array and the alias is broken.
- Rule of thumb: never assume append is "copy-on-write" — it isn't, until it has to be.

## The gotcha

A function that takes a slice and calls `append` on it can silently overwrite the caller's data if the slice has spare capacity. Convention: always return the new slice header (`s = append(s, ...)`) and document whether the function mutates the caller's backing array. Also: a tiny sub-slice of a huge slice keeps the whole backing array alive (GC can't collect it) — copy when you need to forget the parent.

## In code

See the runnable demo in this file. Key output:
- `a: [1 99 3 4 5]` after writing through `b`
- `x: [1 2 999]` after `append` within capacity (alias preserved)
- `p: [1 2 3]` after `append` that exceeds capacity (alias broken)

## Related

- [[slices-basics]] — the (ptr, len, cap) triple
- [[map-iteration-order]] — another "shared backing" surprise
- [[copy-on-write-patterns]] — defensive copying conventions
