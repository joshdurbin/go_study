//go:build ignore

package main

import "fmt"

// singleNumber: XOR fold. x^x = 0 and x^0 = x, so every pair cancels and the
// lone element survives. XOR is commutative and associative — order doesn't matter.
// O(n) time, O(1) space.
func singleNumber(nums []int) int {
	result := 0
	for _, x := range nums {
		result ^= x
	}
	return result
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))          // 1
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))    // 4
	fmt.Println(singleNumber([]int{1}))                // 1
}
