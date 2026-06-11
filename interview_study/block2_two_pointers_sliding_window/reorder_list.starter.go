//go:build ignore

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// reorderList rewires L0→L1→...→Ln in-place to L0→Ln→L1→Ln-1→...
// Three steps: find middle (slow/fast), reverse second half, interleave.
// O(n) time, O(1) space.
func reorderList(head *ListNode) {
	// TODO: implement
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
	a := build(1, 2, 3, 4)
	reorderList(a)
	print(a) // expect 1 4 2 3

	b := build(1, 2, 3, 4, 5)
	reorderList(b)
	print(b) // expect 1 5 2 4 3

	c := build(1)
	reorderList(c)
	print(c) // expect 1
}
