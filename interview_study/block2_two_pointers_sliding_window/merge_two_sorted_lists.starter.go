//go:build ignore

package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

// mergeTwoLists merges two sorted lists by SPLICING NODES (no new allocations).
// Use a dummy head + tail pointer.
func mergeTwoLists(a, b *Node) *Node {
	// TODO: implement
	return nil
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
	print(mergeTwoLists(build(1, 2, 4), build(1, 3, 4))) // expect 1 1 2 3 4 4
	print(mergeTwoLists(nil, build(0)))                  // expect 0
	print(mergeTwoLists(nil, nil))                       // expect (empty)
}
