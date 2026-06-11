//go:build ignore

package main

import (
	"container/heap"
	"fmt"
)

// Heap boilerplate provided — focus on the algorithm.

// MaxHeap: top is largest (holds lower half of stream).
type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{}   { o := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return o }

// MinHeap: top is smallest (holds upper half of stream).
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{}   { o := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return o }

// MedianFinder maintains two heaps. Invariant: |low| - |high| in {0, 1}.
type MedianFinder struct {
	low  *MaxHeap
	high *MinHeap
}

func NewMedianFinder() *MedianFinder {
	return &MedianFinder{low: &MaxHeap{}, high: &MinHeap{}}
}

// AddNum: push to low, sift top to high, rebalance if high outgrew low.
// O(log n).
func (m *MedianFinder) AddNum(n int) {
	// TODO: implement
	_ = heap.Interface(nil)
}

// FindMedian: O(1). If sizes equal, average the two tops; otherwise low's top.
func (m *MedianFinder) FindMedian() float64 {
	// TODO: implement
	return 0
}

func main() {
	mf := NewMedianFinder()
	mf.AddNum(1)
	mf.AddNum(2)
	fmt.Println(mf.FindMedian()) // expect 1.5
	mf.AddNum(3)
	fmt.Println(mf.FindMedian()) // expect 2

	mf2 := NewMedianFinder()
	for _, n := range []int{6, 10, 2, 6, 5, 0} {
		mf2.AddNum(n)
	}
	fmt.Println(mf2.FindMedian()) // expect 5.5
}
