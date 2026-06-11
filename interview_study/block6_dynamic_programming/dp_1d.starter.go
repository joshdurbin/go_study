//go:build ignore

package main

import "fmt"

// rob returns the maximum money you can rob without robbing two adjacent houses.
// Classic 1-D DP: dp[i] = max(dp[i-1], dp[i-2] + nums[i]).
// Optimize to O(1) space if you can.
func rob(nums []int) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))       // expect 4
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))    // expect 12
	fmt.Println(rob([]int{}))                  // expect 0
	fmt.Println(rob([]int{5}))                 // expect 5
}
