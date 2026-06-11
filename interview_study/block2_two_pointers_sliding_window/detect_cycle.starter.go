//go:build ignore

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// detectCycle returns the node where the cycle begins, or nil if no cycle.
// Use Floyd's Tortoise & Hare: phase 1 detects a meeting point inside the cycle,
// phase 2 walks from head and meeting point at equal speed to find the entry.
// O(n) time, O(1) space.
func detectCycle(head *ListNode) *ListNode {
	// TODO: implement
	return nil
}

func build(vals []int, cycleAt int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	var cycleStart *ListNode
	for i, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
		if i == cycleAt {
			cycleStart = cur
		}
	}
	if cycleStart != nil {
		cur.Next = cycleStart
	}
	return dummy.Next
}

func nodeVal(n *ListNode) string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("%d", n.Val)
}

func main() {
	fmt.Println(nodeVal(detectCycle(build([]int{1, 2, 3, 4}, -1)))) // expect nil
	fmt.Println(nodeVal(detectCycle(build([]int{1, 2, 3, 4}, 1))))  // expect 2
	fmt.Println(nodeVal(detectCycle(build([]int{3, 2, 0, -4}, 1)))) // expect 2
}
