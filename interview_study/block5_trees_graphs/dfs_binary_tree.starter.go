//go:build ignore

package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// maxDepth returns the maximum depth (height) of the tree. Use post-order DFS:
// each subtree returns its own depth; the root combines them.
func maxDepth(n *TreeNode) int {
	// TODO: implement
	return 0
}

// isBalanced reports whether the tree is height-balanced (every node's left and
// right subtree heights differ by at most 1). Also post-order.
func isBalanced(n *TreeNode) bool {
	// TODO: implement
	return false
}

func main() {
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(maxDepth(root))   // expect 3
	fmt.Println(isBalanced(root)) // expect true

	skewed := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	fmt.Println(isBalanced(skewed)) // expect false
}
