//go:build ignore

package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// nodeHeap is a min-heap of *ListNode keyed by Val.
type nodeHeap []*ListNode

func (h nodeHeap) Len() int            { return len(h) }
func (h nodeHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h nodeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *nodeHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *nodeHeap) Pop() interface{}   { o := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return o }

// mergeKLists merges k sorted linked lists into one. Push each list head into a
// min-heap; repeatedly pop the smallest and push its next.
// O(n log k) time where n is total nodes, O(k) heap space.
func mergeKLists(lists []*ListNode) *ListNode {
	h := &nodeHeap{}
	for _, head := range lists {
		if head != nil {
			heap.Push(h, head)
		}
	}
	dummy := &ListNode{}
	tail := dummy
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		tail.Next = node
		tail = node
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	return dummy.Next
}

// build constructs a linked list from a slice for testing.
func build(vals []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

// dump returns a slice of values for printing.
func dump(head *ListNode) []int {
	out := []int{}
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}

func main() {
	lists := []*ListNode{
		build([]int{1, 4, 5}),
		build([]int{1, 3, 4}),
		build([]int{2, 6}),
	}
	fmt.Println(dump(mergeKLists(lists))) // [1 1 2 3 4 4 5 6]

	empty := []*ListNode{nil, build([]int{0}), nil}
	fmt.Println(dump(mergeKLists(empty))) // [0]

	fmt.Println(dump(mergeKLists(nil))) // []
}
