//go:build ignore

package main

import "fmt"

// eraseOverlapIntervals returns the minimum number of intervals to remove
// so that the remaining intervals are non-overlapping. Touching is allowed.
func eraseOverlapIntervals(intervals [][]int) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // expect 1
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))         // expect 2
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}}))                 // expect 0
}
