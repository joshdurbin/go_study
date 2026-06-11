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

// minMeetingRooms returns the minimum number of rooms required.
// Sort by start; min-heap of end times. Reuse room if top <= new start.
// O(n log n) time, O(n) space.
func minMeetingRooms(intervals [][]int) int {
	// TODO: implement
	_ = heap.Interface(nil)
	return 0
}

func main() {
	fmt.Println(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}})) // expect 2
	fmt.Println(minMeetingRooms([][]int{{7, 10}, {2, 4}}))            // expect 1
	fmt.Println(minMeetingRooms([][]int{{1, 5}, {2, 6}, {3, 7}}))     // expect 3
	fmt.Println(minMeetingRooms([][]int{}))                           // expect 0
}
