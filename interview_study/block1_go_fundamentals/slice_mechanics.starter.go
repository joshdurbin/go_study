//go:build ignore

package main

import "fmt"

// removeDuplicates removes duplicates from a in-place, preserving original order,
// and returns the trimmed slice. Do not allocate a second slice for the result.
// Target O(n) time using a write-head two-pointer.
func removeDuplicates(a []int) []int {
	// TODO: implement
	return a
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 2, 2, 3, 1, 4})) // expect [1 2 3 4]
	fmt.Println(removeDuplicates([]int{5, 5, 5}))          // expect [5]
}
