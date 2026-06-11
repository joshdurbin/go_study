//go:build ignore

package main

import (
	"fmt"
	"sort"
)

// eraseOverlapIntervals: greedy — sort by end time and keep the interval whose
// start is >= last kept end. Every skipped interval is one removal.
// O(n log n) time, O(1) space (excluding sort scratch).
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })

	removals := 0
	end := intervals[0][1]
	for _, x := range intervals[1:] {
		if x[0] >= end {
			end = x[1]
		} else {
			removals++
		}
	}
	return removals
}

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // 1
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))         // 2
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}}))                 // 0
}
