//go:build ignore

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle reports whether the linked list contains a cycle.
// Use Floyd's tortoise-and-hare (fast/slow pointers). O(n) time, O(1) space.
func hasCycle(head *ListNode) bool {
	// TODO: implement
	return false
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

func main() {
	fmt.Println(hasCycle(build([]int{1, 2, 3, 4}, -1))) // expect false
	fmt.Println(hasCycle(build([]int{1, 2, 3, 4}, 1)))  // expect true
}
