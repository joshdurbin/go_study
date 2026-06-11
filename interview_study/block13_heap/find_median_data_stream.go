//go:build ignore

package main

import (
	"container/heap"
	"fmt"
)

// MaxHeap of ints — top is largest. Holds the LOWER half of the stream.
type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{}   { o := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return o }

// MinHeap of ints — top is smallest. Holds the UPPER half of the stream.
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{}   { o := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return o }

// MedianFinder uses two heaps. Invariant: low.Len() == high.Len() or
// low.Len() == high.Len()+1. Max of low <= min of high.
type MedianFinder struct {
	low  *MaxHeap // lower half
	high *MinHeap // upper half
}

func NewMedianFinder() *MedianFinder {
	return &MedianFinder{low: &MaxHeap{}, high: &MinHeap{}}
}

// AddNum: push to low, then move low's top to high, then rebalance if high
// outgrew low. O(log n) per insertion.
func (m *MedianFinder) AddNum(n int) {
	heap.Push(m.low, n)
	heap.Push(m.high, heap.Pop(m.low))
	if m.high.Len() > m.low.Len() {
		heap.Push(m.low, heap.Pop(m.high))
	}
}

// FindMedian: O(1). If sizes equal, average the two tops; otherwise low's top.
func (m *MedianFinder) FindMedian() float64 {
	if m.low.Len() > m.high.Len() {
		return float64((*m.low)[0])
	}
	return (float64((*m.low)[0]) + float64((*m.high)[0])) / 2.0
}

func main() {
	mf := NewMedianFinder()
	mf.AddNum(1)
	mf.AddNum(2)
	fmt.Println(mf.FindMedian()) // 1.5
	mf.AddNum(3)
	fmt.Println(mf.FindMedian()) // 2

	mf2 := NewMedianFinder()
	for _, n := range []int{6, 10, 2, 6, 5, 0} {
		mf2.AddNum(n)
	}
	fmt.Println(mf2.FindMedian()) // 5.5
}
