## Hint 1
Post-order DFS. Each subtree returns "I contain p or q (or am one of them)" or nil.

```go
func lca(root, p, q *TreeNode) *TreeNode {
    if root == nil || root == p || root == q {
        return root
    }
    // recurse left and right
}
```

## Hint 2
Recurse both sides, then check the two results.

```go
l := lca(root.Left, p, q)
r := lca(root.Right, p, q)
```

## Hint 3
Both non-nil → THIS node is the LCA. Otherwise propagate whichever side was non-nil.

```go
if l != nil && r != nil { return root }
if l != nil { return l }
return r
```
