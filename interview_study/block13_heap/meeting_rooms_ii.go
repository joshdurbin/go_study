//go:build ignore

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// MinHeap of ints — tracks end times of in-use rooms; top is earliest ending.
type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{}   { o := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return o }

// minMeetingRooms: sort by start; min-heap of end times. For each meeting, if
// the earliest-ending room is free (end <= start), reuse it (pop). Always
// push the new end. Final heap size = peak concurrent meetings.
// O(n log n) time, O(n) space.
func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	h := &MinHeap{}
	for _, m := range intervals {
		if h.Len() > 0 && (*h)[0] <= m[0] {
			heap.Pop(h) // reuse the room ending soonest
		}
		heap.Push(h, m[1])
	}
	return h.Len()
}

func main() {
	fmt.Println(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}})) // 2
	fmt.Println(minMeetingRooms([][]int{{7, 10}, {2, 4}}))            // 1
	fmt.Println(minMeetingRooms([][]int{{1, 5}, {2, 6}, {3, 7}}))     // 3
	fmt.Println(minMeetingRooms([][]int{}))                           // 0
}
