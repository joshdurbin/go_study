//go:build ignore

package main

import "fmt"

// twoSum returns the indices of two numbers in nums that add up to target.
// Assume exactly one solution exists; you may not use the same element twice.
func twoSum(nums []int, target int) []int {
	// TODO: implement
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // expect [0 1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // expect [1 2]
	fmt.Println(twoSum([]int{3, 3}, 6))         // expect [0 1]
}
