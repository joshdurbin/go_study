//go:build ignore

package main

import "fmt"

// maxSubArray: Kadane's algorithm. cur = best sum ending at i; best = global max.
// O(n) time, O(1) space.
func maxSubArray(nums []int) int {
	cur, best := nums[0], nums[0]
	for _, x := range nums[1:] {
		if cur+x < x {
			cur = x // restart at x
		} else {
			cur += x // extend the run
		}
		if cur > best {
			best = cur
		}
	}
	return best
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(maxSubArray([]int{1}))                             // 1
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))                // 23
}
