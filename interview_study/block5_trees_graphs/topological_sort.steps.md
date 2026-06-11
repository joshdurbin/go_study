## Hint 1
Kahn's algorithm: compute in-degrees, start from all 0-in-degree nodes.

```go
inDeg := make([]int, n)
for u := range adj {
    for _, v := range adj[u] { inDeg[v]++ }
}
q := []int{}
for u, d := range inDeg {
    if d == 0 { q = append(q, u) }
}
```

## Hint 2
Pop one at a time, emit to result, decrement each neighbor's in-degree, enqueue any new zeros.

```go
var order []int
for len(q) > 0 {
    u := q[0]; q = q[1:]
    order = append(order, u)
    for _, v := range adj[u] {
        inDeg[v]--
        if inDeg[v] == 0 { q = append(q, v) }
    }
}
```

## Hint 3
If `len(order) < n`, the graph has a cycle — no valid order exists. This dual purpose makes Kahn's the go-to for "course schedule" problems.

```go
if len(order) != n {
    return nil // cycle detected
}
return order
```
