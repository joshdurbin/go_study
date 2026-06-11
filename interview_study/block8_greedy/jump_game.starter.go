//go:build ignore

package main

import "fmt"

// canJump reports whether you can reach the last index of nums starting from
// index 0, where nums[i] is the maximum jump length from index i.
func canJump(nums []int) bool {
	// TODO: implement
	return false
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4})) // expect true
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) // expect false
	fmt.Println(canJump([]int{0}))             // expect true
	fmt.Println(canJump([]int{2, 0, 0}))       // expect true
}
