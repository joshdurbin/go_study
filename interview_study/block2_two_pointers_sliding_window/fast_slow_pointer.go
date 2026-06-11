//go:build ignore

package main

import "fmt"

// findDuplicate uses Floyd's cycle detection.
// Phase 1: find meeting point inside the cycle.
// Phase 2: find cycle entry (the duplicate).
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2})) // 2
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2})) // 3
}
