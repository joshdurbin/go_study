# Map iteration order

## The question

Why does `for k := range m` give a different order every time? How do I get deterministic iteration?

## The answer

- The Go runtime **intentionally** randomizes map iteration order. The starting bucket is chosen with a per-range randomized offset.
- This was added in Go 1.0 specifically to prevent code from depending on insertion order, hash order, or any other "accidental" ordering. If iteration were deterministic, real programs would silently lock in that order — and break later when the runtime improves.
- For determinism, collect the keys into a slice and sort: `keys := slices.Sorted(maps.Keys(m))` (Go 1.23+) or the classic `sort.Strings(keys)`.
- For "insertion-order" semantics, Go has no built-in — use a separate slice as an order log, or a third-party `orderedmap` library.

## The gotcha

Two ranges over the same map in the same program produce different orders. Test assertions like `assert.Equal(expected, fmt.Sprint(m))` flake. JSON-encoding a map of maps gives non-reproducible bytes (`encoding/json` does sort top-level keys, but that's a property of the encoder, not the range).

## In code

See the runnable demo in this file. Key output:
- `range 1: ...` and `range 2: ...` — orders differ
- `sorted: alpha=1 bravo=2 charlie=3 delta=4 echo=5`

## Related

- [[slice-aliasing-footguns]] — another "Go intentionally surprises you" case
- [[json-encoding-determinism]] — `encoding/json` sorts map keys; protobuf may not
