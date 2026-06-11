## Hint 1
Dijkstra = BFS + a min-heap keyed by tentative distance. Initialize all distances to ∞ except src.

```go
dist := make([]int, n)
for i := range dist { dist[i] = math.MaxInt }
dist[src] = 0
h := &pq{{node: src, dist: 0}}
```

## Hint 2
Pop the closest unfinished node, relax outgoing edges.

```go
for h.Len() > 0 {
    cur := heap.Pop(h).(item)
    // skip stale, then relax neighbors
}
```

## Hint 3
container/heap has no decrease-key, so the same node may be pushed multiple times. Skip stale entries by comparing distances.

```go
if cur.dist > dist[cur.node] { continue }
for _, e := range adj[cur.node] {
    nd := cur.dist + e.w
    if nd < dist[e.to] {
        dist[e.to] = nd
        heap.Push(h, item{node: e.to, dist: nd})
    }
}
```
