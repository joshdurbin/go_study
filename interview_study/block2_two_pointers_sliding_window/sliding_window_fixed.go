//go:build ignore

package main

import "fmt"

func maxSumK(nums []int, k int) int {
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	maxSum := windowSum
	for i := k; i < len(nums); i++ {
		windowSum += nums[i] - nums[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return maxSum
}

func main() {
	fmt.Println(maxSumK([]int{2, 1, 5, 1, 3, 2}, 3)) // 9
	fmt.Println(maxSumK([]int{2, 3, 4, 1, 5}, 2))    // 7
}
