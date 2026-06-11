//go:build ignore

package main

import "fmt"

// subsets returns all 2^n subsets of nums (distinct ints). Order of subsets
// doesn't matter. Use the choose/un-choose backtracking template with a start
// index to avoid duplicates.
func subsets(nums []int) [][]int {
	// TODO: implement
	return nil
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3})) // expect [[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]
	fmt.Println(subsets([]int{0}))       // expect [[] [0]]
}
