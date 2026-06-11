//go:build ignore

package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

// mergeTwoLists: dummy head + tail pointer. Splice the smaller node each step.
// O(m+n) time, O(1) extra space.
func mergeTwoLists(a, b *Node) *Node {
	dummy := &Node{}
	tail := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Next = a
			a = a.Next
		} else {
			tail.Next = b
			b = b.Next
		}
		tail = tail.Next
	}
	if a != nil {
		tail.Next = a
	} else {
		tail.Next = b
	}
	return dummy.Next
}

func build(vals ...int) *Node {
	d := &Node{}
	cur := d
	for _, v := range vals {
		cur.Next = &Node{Val: v}
		cur = cur.Next
	}
	return d.Next
}

func print(h *Node) {
	for h != nil {
		fmt.Printf("%d ", h.Val)
		h = h.Next
	}
	fmt.Println()
}

func main() {
	print(mergeTwoLists(build(1, 2, 4), build(1, 3, 4))) // 1 1 2 3 4 4
	print(mergeTwoLists(nil, build(0)))                  // 0
	print(mergeTwoLists(nil, nil))                       // (empty)
}
