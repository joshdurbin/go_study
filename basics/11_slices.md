# Slices

The most-used and most-misunderstood Go data structure. A slice is a (ptr, len, cap) triple over a backing array.

## Worth knowing

- `make([]T, len, cap)` pre-allocates capacity. `make([]T, 0, expectedSize)` avoids reallocations in `append`.
- `append` returns a new slice header — always assign it: `s = append(s, x)`.
- When cap is exceeded, append allocates a **new** backing array. Pre-allocate when you can.
- Slicing shares memory: `b := a[1:4]` — modifying `b[0]` also modifies `a[1]`.
- Pre-allocate with `make`, copy with `copy(dst, src)`. Don't try to use append to copy — semantics are subtle.

## Common gotchas

- **Hidden retention**: `small := bigSlice[5:10]` keeps the entire `bigSlice` backing array alive (GC can't free it). Copy if you need a tiny excerpt of a huge slice.
- **Aliasing bug**: passing a slice and later appending may or may not affect the caller depending on cap. Convention: a function that grows a slice should return the new slice.

## Interview-grade

Be ready to explain len vs cap, when append re-allocates, and the aliasing semantics. These come up constantly.
