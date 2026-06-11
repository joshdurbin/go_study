//go:build ignore

package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// serialize encodes the tree as a string. Use pre-order DFS with a sentinel
// (e.g., "#") for nil children, separated by commas.
func serialize(root *TreeNode) string {
	// TODO: implement
	return ""
}

// deserialize reconstructs the tree from serialize's output.
func deserialize(s string) *TreeNode {
	// TODO: implement
	return nil
}

func main() {
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		}}
	s := serialize(root)
	fmt.Println(s)
	// round-trip should serialize identically
	fmt.Println(serialize(deserialize(s)))
}
