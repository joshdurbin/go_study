//go:build ignore

package main

import "fmt"

// maxSubArray returns the largest sum among all contiguous subarrays of nums.
// Target O(n) time, O(1) space (Kadane's algorithm).
func maxSubArray(nums []int) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // expect 6
	fmt.Println(maxSubArray([]int{1}))                             // expect 1
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))                // expect 23
}
