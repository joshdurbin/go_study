//go:build ignore

package main

import (
	"fmt"
	"sort"
)

// merge: sort by start, then sweep. If next.start <= cur.end, extend cur.end.
// O(n log n) time dominated by the sort.
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	out := [][]int{intervals[0]}
	for _, x := range intervals[1:] {
		last := out[len(out)-1]
		if x[0] <= last[1] {
			if x[1] > last[1] {
				last[1] = x[1]
			}
		} else {
			out = append(out, x)
		}
	}
	return out
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // [[1 6] [8 10] [15 18]]
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))                    // [[1 5]]
}
