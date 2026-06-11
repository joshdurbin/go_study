## Hint 1
For tree property questions (height, diameter, balanced check), use POST-order — children's answers feed the parent's decision.

```go
var dfs func(*TreeNode) int
dfs = func(n *TreeNode) int {
    if n == nil { return 0 }
    // recurse left & right first, then combine
}
```

## Hint 2
Combine the two subtree answers into the parent's return value.

```go
l := dfs(n.Left)
r := dfs(n.Right)
return 1 + max(l, r) // example: height
```

## Hint 3
Pre-order is for "decide at the root and pass info down" (DFS path collection). In-order on a BST visits values in sorted order — useful by itself.

```go
// in-order: prints BST values ascending
var inorder func(*TreeNode)
inorder = func(n *TreeNode) {
    if n == nil { return }
    inorder(n.Left)
    fmt.Println(n.Val)
    inorder(n.Right)
}
```
