//go:build ignore

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// reorderList rewires L0→L1→...→Ln into L0→Ln→L1→Ln-1→...
// Three sub-routines: find middle (slow/fast), reverse second half, interleave.
// O(n) time, O(1) space.
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// 1. find middle: slow ends at the last node of the first half
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 2. reverse the second half
	var prev *ListNode
	cur := slow.Next
	slow.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	// 3. interleave the two halves
	a, b := head, prev
	for b != nil {
		aNext, bNext := a.Next, b.Next
		a.Next = b
		b.Next = aNext
		a, b = aNext, bNext
	}
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
	print(a) // 1 4 2 3

	b := build(1, 2, 3, 4, 5)
	reorderList(b)
	print(b) // 1 5 2 4 3

	c := build(1)
	reorderList(c)
	print(c) // 1
}
