# Maps

Hash tables built into the language. `map[K]V` with any comparable K.

## Worth knowing

- Always `make` before use: `m := make(map[string]int)`. The zero-value map is nil — reading from it works (returns the zero value of V), but writing panics.
- Lookup returns `(value, ok)`: `v, ok := m[k]`. The bare form `m[k]` returns the zero value on miss — sometimes you need to distinguish.
- `delete(m, k)` removes a key. Removing a non-existent key is a no-op, not an error.
- Iteration order is **randomized** (intentionally — Go randomizes it to prevent code from depending on it).
- Maps are **not** safe for concurrent use. Use `sync.Map` or a mutex.

## Idioms

- Counter: `counts[word]++` (works on a missing key thanks to zero values).
- Set: `map[T]struct{}` (zero bytes per entry, vs. 1 byte for `map[T]bool`).
- Group: `groups[key] = append(groups[key], value)`.

## Interview-grade

The randomized iteration order is a frequent question. So is "what happens on concurrent map writes" (the runtime panics with "concurrent map writes").
