//go:build ignore

package main

import "fmt"

// productExceptSelf returns out where out[i] is the product of every element
// of nums except nums[i]. You may NOT use division. Target O(n) time.
func productExceptSelf(nums []int) []int {
	// TODO: implement
	return nil
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))         // expect [24 12 8 6]
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))    // expect [0 0 9 0 0]
}
