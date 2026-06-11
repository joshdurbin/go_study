//go:build ignore

package main

import "fmt"

// merge returns a sorted slice of intervals with all overlapping intervals merged.
// Each interval is [start, end]. Touching intervals (next.start == prev.end) overlap.
func merge(intervals [][]int) [][]int {
	// TODO: implement
	return nil
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // expect [[1 6] [8 10] [15 18]]
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))                    // expect [[1 5]]
	fmt.Println(merge([][]int{{1, 4}, {0, 4}}))                    // expect [[0 4]]
}
