//go:build ignore

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// MinHeap implements heap.Interface for ints (min at top).
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// kSmallest returns the K smallest elements in sorted order.
func kSmallest(nums []int, k int) []int {
	h := MinHeap(append([]int{}, nums...))
	heap.Init(&h)
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(&h).(int)
	}
	return result
}

func main() {
	nums := []int{5, 3, 8, 1, 9, 2, 7}
	result := kSmallest(nums, 3)
	sort.Ints(result)
	fmt.Println(result) // [1 2 3]
}
