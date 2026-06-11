//go:build ignore

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// serialize: preorder DFS, "#" for nil. O(n) time.
func serialize(root *TreeNode) string {
	var sb strings.Builder
	var walk func(*TreeNode)
	walk = func(n *TreeNode) {
		if n == nil {
			sb.WriteString("#,")
			return
		}
		sb.WriteString(strconv.Itoa(n.Val))
		sb.WriteByte(',')
		walk(n.Left)
		walk(n.Right)
	}
	walk(root)
	return strings.TrimSuffix(sb.String(), ",")
}

// deserialize: consume tokens in the same preorder, recursing on each.
func deserialize(s string) *TreeNode {
	if s == "" {
		return nil
	}
	tokens := strings.Split(s, ",")
	i := 0
	var build func() *TreeNode
	build = func() *TreeNode {
		if i >= len(tokens) {
			return nil
		}
		t := tokens[i]
		i++
		if t == "#" {
			return nil
		}
		v, _ := strconv.Atoi(t)
		n := &TreeNode{Val: v}
		n.Left = build()
		n.Right = build()
		return n
	}
	return build()
}

func main() {
	root := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		}}
	s := serialize(root)
	fmt.Println(s)
	back := deserialize(s)
	fmt.Println(serialize(back)) // round-trip
}
