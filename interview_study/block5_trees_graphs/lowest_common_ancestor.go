//go:build ignore

package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// lca: post-order DFS. If a subtree contains p or q (or is p/q itself), return
// that finding upward. The first node where BOTH children return non-nil is
// the LCA. O(n) time, O(h) stack.
func lca(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	l := lca(root.Left, p, q)
	r := lca(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l != nil {
		return l
	}
	return r
}

func main() {
	n7, n4 := &TreeNode{Val: 7}, &TreeNode{Val: 4}
	n2 := &TreeNode{Val: 2, Left: n7, Right: n4}
	n6 := &TreeNode{Val: 6}
	n5 := &TreeNode{Val: 5, Left: n6, Right: n2}
	n0, n8 := &TreeNode{Val: 0}, &TreeNode{Val: 8}
	n1 := &TreeNode{Val: 1, Left: n0, Right: n8}
	root := &TreeNode{Val: 3, Left: n5, Right: n1}

	fmt.Println(lca(root, n5, n1).Val) // 3
	fmt.Println(lca(root, n5, n4).Val) // 5
	fmt.Println(lca(root, n7, n4).Val) // 2
}
