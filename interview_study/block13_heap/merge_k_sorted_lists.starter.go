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

// Heap boilerplate provided — focus on the algorithm.
type nodeHeap []*ListNode

func (h nodeHeap) Len() int            { return len(h) }
func (h nodeHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h nodeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *nodeHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *nodeHeap) Pop() interface{}   { o := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return o }

// mergeKLists merges k sorted lists. Use the nodeHeap above.
// Push each non-nil head, then repeatedly pop smallest and push popped.Next.
// O(n log k) time, O(k) space.
func mergeKLists(lists []*ListNode) *ListNode {
	// TODO: implement
	_ = heap.Interface(nil)
	return nil
}

func build(vals []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

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
	fmt.Println(dump(mergeKLists(lists))) // expect [1 1 2 3 4 4 5 6]

	empty := []*ListNode{nil, build([]int{0}), nil}
	fmt.Println(dump(mergeKLists(empty))) // expect [0]

	fmt.Println(dump(mergeKLists(nil))) // expect []
}
