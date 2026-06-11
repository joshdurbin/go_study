//go:build ignore

package main

import "fmt"

// nextGreaterElements treats nums as circular and returns, for each index, the
// next greater value when scanning forward (wrapping). -1 if none.
// Hint: monotonic decreasing stack of indices, iterate 2n times with i % n.
func nextGreaterElements(nums []int) []int {
	// TODO: implement
	return nil
}

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))       // expect [2 -1 2]
	fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 3})) // expect [2 3 4 -1 4]
	fmt.Println(nextGreaterElements([]int{5, 4, 3, 2, 1})) // expect [-1 5 5 5 5]
}
