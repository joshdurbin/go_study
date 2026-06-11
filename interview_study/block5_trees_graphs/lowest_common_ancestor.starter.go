//go:build ignore

package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// lca returns the lowest common ancestor of p and q in the binary tree rooted at root.
// Both p and q are guaranteed to be present in the tree.
// Use post-order DFS: each subtree reports back whether it contains p or q.
func lca(root, p, q *TreeNode) *TreeNode {
	// TODO: implement
	return nil
}

func main() {
	n7, n4 := &TreeNode{Val: 7}, &TreeNode{Val: 4}
	n2 := &TreeNode{Val: 2, Left: n7, Right: n4}
	n6 := &TreeNode{Val: 6}
	n5 := &TreeNode{Val: 5, Left: n6, Right: n2}
	n0, n8 := &TreeNode{Val: 0}, &TreeNode{Val: 8}
	n1 := &TreeNode{Val: 1, Left: n0, Right: n8}
	root := &TreeNode{Val: 3, Left: n5, Right: n1}

	fmt.Println(lca(root, n5, n1).Val) // expect 3
	fmt.Println(lca(root, n5, n4).Val) // expect 5
	fmt.Println(lca(root, n7, n4).Val) // expect 2
}
