//go:build ignore

package main

import "fmt"

// insert places newInterval into a sorted, non-overlapping list of intervals,
// merging any that overlap. Result must remain sorted and non-overlapping.
func insert(intervals [][]int, newInterval []int) [][]int {
	// TODO: implement
	return nil
}

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))                          // expect [[1 5] [6 9]]
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})) // expect [[1 2] [3 10] [12 16]]
	fmt.Println(insert([][]int{}, []int{5, 7}))                                        // expect [[5 7]]
}
