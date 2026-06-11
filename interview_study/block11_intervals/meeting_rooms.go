//go:build ignore

package main

import (
	"fmt"
	"sort"
)

// canAttendMeetings: sort by start, then any meeting starting before the
// previous one ended is a conflict.
// O(n log n) time, O(1) space (excluding sort scratch).
func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canAttendMeetings([][]int{{0, 30}, {5, 10}, {15, 20}})) // false
	fmt.Println(canAttendMeetings([][]int{{7, 10}, {2, 4}}))            // true
	fmt.Println(canAttendMeetings([][]int{{1, 5}, {5, 8}}))             // true
}
