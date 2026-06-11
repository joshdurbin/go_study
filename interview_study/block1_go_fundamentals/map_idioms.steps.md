## Hint 1
Three idioms to know: comma-ok lookup, `m[k]++` (zero-value friendly), and `delete(m, k)`. None allocate on a missing key.

```go
counts := make(map[string]int)
for _, w := range words {
    counts[w]++ // works even if w not in map yet
}
```

## Hint 2
Comma-ok is the ONLY way to distinguish "absent" from "present with zero value".

```go
if v, ok := m[k]; ok {
    fmt.Println("present:", v)
} else {
    fmt.Println("missing")
}
```

## Hint 3
For sets, prefer `map[T]struct{}` over `map[T]bool` — zero bytes per entry. Iteration order is randomized; sort keys if you need stability.

```go
set := make(map[string]struct{})
set["a"] = struct{}{}
if _, ok := set["a"]; ok { /* member */ }
```
