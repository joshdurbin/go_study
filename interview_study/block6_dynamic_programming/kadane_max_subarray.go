//go:build ignore

package main

import "fmt"

// maxSubArray: Kadane's algorithm as a DP.
// State:   dp[i] = max sum of a contiguous subarray ending exactly at index i.
// Base:    dp[0] = nums[0]
// Recur:   dp[i] = max(nums[i], dp[i-1] + nums[i])
// Answer:  max(dp[0..n-1])
// Only dp[i-1] is read, so collapse the table into a single rolling scalar.
// O(n) time, O(1) space.
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0 // documented edge case; classic LeetCode framing assumes non-empty
	}
	curr, best := nums[0], nums[0]
	for _, x := range nums[1:] {
		// dp[i] = max(x, dp[i-1] + x)
		if curr+x > x {
			curr += x
		} else {
			curr = x
		}
		if curr > best {
			best = curr
		}
	}
	return best
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(maxSubArray([]int{1}))                             // 1
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))                // 23
	fmt.Println(maxSubArray([]int{-3, -1, -4, -2}))                // -1
	fmt.Println(maxSubArray([]int{}))                              // 0
}
