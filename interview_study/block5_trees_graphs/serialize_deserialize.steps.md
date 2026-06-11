## Hint 1
Pre-order DFS with explicit nil markers. The marker is what makes the tree structure recoverable.

```go
func serialize(root *TreeNode) string {
    var sb strings.Builder
    var walk func(*TreeNode)
    walk = func(n *TreeNode) {
        if n == nil { sb.WriteString("#,"); return }
        // write value, recurse L, recurse R
    }
    walk(root)
    return strings.TrimSuffix(sb.String(), ",")
}
```

## Hint 2
Body of walk for non-nil nodes:

```go
sb.WriteString(strconv.Itoa(n.Val))
sb.WriteByte(',')
walk(n.Left)
walk(n.Right)
```

## Hint 3
Deserialize: split tokens, consume in the same pre-order. Share an index pointer so both children advance through the SAME stream.

```go
tokens := strings.Split(s, ",")
i := 0
var build func() *TreeNode
build = func() *TreeNode {
    t := tokens[i]; i++
    if t == "#" { return nil }
    v, _ := strconv.Atoi(t)
    n := &TreeNode{Val: v}
    n.Left  = build()
    n.Right = build()
    return n
}
return build()
```
