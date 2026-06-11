//go:build ignore

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList reverses a singly-linked list in-place and returns the new head.
// O(n) time, O(1) space.
func reverseList(head *ListNode) *ListNode {
	// TODO: implement
	return nil
}

func build(vals ...int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func print(h *ListNode) {
	for h != nil {
		fmt.Printf("%d ", h.Val)
		h = h.Next
	}
	fmt.Println()
}

func main() {
	print(reverseList(build(1, 2, 3, 4, 5))) // expect 5 4 3 2 1
	print(reverseList(nil))                  // expect (empty line)
}
