//go:build ignore

package main

import "fmt"

// twoSumSorted returns the 1-indexed positions of two numbers in a SORTED
// array that add up to target. Assume exactly one solution.
func twoSumSorted(a []int, target int) []int {
	// TODO: implement
	return nil
}

func main() {
	fmt.Println(twoSumSorted([]int{2, 7, 11, 15}, 9)) // expect [1 2]
	fmt.Println(twoSumSorted([]int{1, 2, 3, 4}, 7))   // expect [3 4]
}
