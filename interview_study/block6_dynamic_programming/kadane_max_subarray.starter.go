//go:build ignore

package main

import "fmt"

// maxSubArray returns the largest sum among all contiguous subarrays of nums.
// DP framing: dp[i] = max(nums[i], dp[i-1] + nums[i]); answer = max(dp...).
// Collapse to O(1) space with a rolling scalar. Watch the all-negative case.
func maxSubArray(nums []int) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // expect 6
	fmt.Println(maxSubArray([]int{1}))                             // expect 1
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))                // expect 23
	fmt.Println(maxSubArray([]int{-3, -1, -4, -2}))                // expect -1
	fmt.Println(maxSubArray([]int{}))                              // expect 0
}
