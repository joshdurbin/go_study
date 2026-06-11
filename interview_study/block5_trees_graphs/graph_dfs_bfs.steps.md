## Hint 1
Graphs allow cycles, so always carry a `visited` set. Adjacency list is the standard representation.

```go
adj := make(map[int][]int) // node → neighbors
visited := make(map[int]bool)
```

## Hint 2
DFS with recursion + visited check at entry.

```go
var dfs func(int)
dfs = func(u int) {
    if visited[u] { return }
    visited[u] = true
    for _, v := range adj[u] {
        dfs(v)
    }
}
```

## Hint 3
BFS with queue + mark visited at ENQUEUE time (not dequeue) — otherwise you enqueue the same node multiple times.

```go
q := []int{src}
visited[src] = true
for len(q) > 0 {
    u := q[0]; q = q[1:]
    for _, v := range adj[u] {
        if !visited[v] {
            visited[v] = true
            q = append(q, v)
        }
    }
}
```
