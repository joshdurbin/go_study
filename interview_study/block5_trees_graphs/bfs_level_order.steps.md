## Hint 1
Level-by-level traversal: snapshot the queue size at the start of each level — that's exactly how many nodes belong to that level.

```go
if root == nil { return nil }
q := []*TreeNode{root}
var out [][]int
for len(q) > 0 {
    size := len(q)
    // process `size` nodes — that's this level
}
```

## Hint 2
Inside the level loop, dequeue once per node, capture its value, enqueue its children.

```go
level := make([]int, 0, size)
for i := 0; i < size; i++ {
    n := q[0]; q = q[1:]
    level = append(level, n.Val)
    if n.Left  != nil { q = append(q, n.Left)  }
    if n.Right != nil { q = append(q, n.Right) }
}
out = append(out, level)
```

## Hint 3
Variations share the skeleton: right-side view (take only the last node per level), zigzag (reverse alternate levels), level averages.

```go
// right-side view:
out = append(out, q[size-1].Val) // last node of the level
```
