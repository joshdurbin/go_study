//go:build ignore

package main

import "fmt"

// singleNumber returns the one element that appears exactly once in nums.
// Every other element appears exactly twice. O(n) time, O(1) space.
func singleNumber(nums []int) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))       // expect 1
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2})) // expect 4
	fmt.Println(singleNumber([]int{1}))             // expect 1
}
