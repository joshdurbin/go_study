//go:build ignore

package main

import (
	"container/heap"
	"fmt"
)

// Heap boilerplate provided — focus on the algorithm.
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{}   { o := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return o }

// findKthLargest returns the k-th largest element using a min-heap of size k.
// Push each value; if heap grows past k, pop the smallest. Top is the answer.
// O(n log k) time, O(k) space.
func findKthLargest(nums []int, k int) int {
	// TODO: implement
	_ = heap.Interface(nil)
	return 0
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))          // expect 5
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)) // expect 4
	fmt.Println(findKthLargest([]int{1}, 1))                         // expect 1
}
