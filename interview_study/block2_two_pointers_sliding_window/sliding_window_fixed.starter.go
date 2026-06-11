//go:build ignore

package main

import "fmt"

// maxSumFixedWindow returns the maximum sum among contiguous subarrays of length k.
// Maintain the sum incrementally — do not recompute the window from scratch.
func maxSumFixedWindow(a []int, k int) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println(maxSumFixedWindow([]int{2, 1, 5, 1, 3, 2}, 3)) // expect 9 (5+1+3)
	fmt.Println(maxSumFixedWindow([]int{-1, -2, -3}, 2))       // expect -3
}
