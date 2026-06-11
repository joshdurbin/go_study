//go:build ignore

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ── Part A: Max Depth ───────────────────────────────────────────────────────

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := maxDepth(root.Left), maxDepth(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}

// maxDepthIterative uses an explicit stack storing (node, depth) pairs.
func maxDepthIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}
	type frame struct {
		node  *TreeNode
		depth int
	}
	stack := []frame{{root, 1}}
	best := 0
	for len(stack) > 0 {
		f := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if f.depth > best {
			best = f.depth
		}
		if f.node.Left != nil {
			stack = append(stack, frame{f.node.Left, f.depth + 1})
		}
		if f.node.Right != nil {
			stack = append(stack, frame{f.node.Right, f.depth + 1})
		}
	}
	return best
}

// ── Part B: Has Path Sum ────────────────────────────────────────────────────

func hasPathSum(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == target
	}
	rem := target - root.Val
	return hasPathSum(root.Left, rem) || hasPathSum(root.Right, rem)
}

func main() {
	root := &TreeNode{5,
		&TreeNode{4, &TreeNode{11,
			&TreeNode{7, nil, nil},
			&TreeNode{2, nil, nil}}, nil},
		&TreeNode{8,
			&TreeNode{13, nil, nil},
			&TreeNode{4, nil, &TreeNode{1, nil, nil}}},
	}
	fmt.Println(maxDepth(root))          // 4
	fmt.Println(maxDepthIterative(root)) // 4
	fmt.Println(hasPathSum(root, 22))    // true
	fmt.Println(hasPathSum(root, 26))    // false
}
