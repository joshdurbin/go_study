//go:build ignore

package main

import "fmt"

// longestSubarraySumAtMost returns the length of the longest contiguous
// subarray whose sum is at most maxSum. Assume non-negative numbers (so the
// sum is monotonic in window length).
func longestSubarraySumAtMost(a []int, maxSum int) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println(longestSubarraySumAtMost([]int{1, 2, 3, 1, 1}, 4)) // expect 3 ([3,1,1] sums to 5? no — try [2,1,1])
	fmt.Println(longestSubarraySumAtMost([]int{1, 1, 1, 1}, 2))    // expect 2
	fmt.Println(longestSubarraySumAtMost([]int{5, 1}, 0))          // expect 0
}
