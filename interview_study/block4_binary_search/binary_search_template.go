//go:build ignore

package main

import "fmt"

func searchRotated(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}
		// left half is sorted
		if nums[lo] <= nums[mid] {
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			// right half is sorted
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(searchRotated([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4
	fmt.Println(searchRotated([]int{4, 5, 6, 7, 0, 1, 2}, 3)) // -1
	fmt.Println(searchRotated([]int{1}, 0))                    // -1
}
