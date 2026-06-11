//go:build ignore

package main

import "fmt"

// insert: three-phase walk. Append intervals strictly before newInterval, merge
// overlapping ones into newInterval, then append the rest.
// O(n) time, O(n) space.
func insert(intervals [][]int, newInterval []int) [][]int {
	out := make([][]int, 0, len(intervals)+1)
	i, n := 0, len(intervals)

	// Phase 1: strictly before — intervals that end before newInterval starts.
	for i < n && intervals[i][1] < newInterval[0] {
		out = append(out, intervals[i])
		i++
	}
	// Phase 2: overlapping — fold them into newInterval.
	for i < n && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	out = append(out, newInterval)
	// Phase 3: strictly after.
	for i < n {
		out = append(out, intervals[i])
		i++
	}
	return out
}

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))                          // [[1 5] [6 9]]
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})) // [[1 2] [3 10] [12 16]]
	fmt.Println(insert([][]int{}, []int{5, 7}))                                        // [[5 7]]
}
