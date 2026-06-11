//go:build ignore

package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// levelOrder returns a slice of levels, each containing the node values at that depth.
// Use BFS with a queue. Snapshot len(queue) at the start of each level.
func levelOrder(root *TreeNode) [][]int {
	// TODO: implement
	return nil
}

func main() {
	root := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	fmt.Println(levelOrder(root)) // expect [[3] [9 20] [15 7]]
}
