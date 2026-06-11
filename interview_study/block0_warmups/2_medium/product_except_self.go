//go:build ignore

package main

import "fmt"

// productExceptSelf: two passes.
// Pass 1: out[i] = product of everything LEFT of i.
// Pass 2: walk right-to-left, multiplying by a running suffix product.
// O(n) time, O(1) extra space (output array doesn't count by convention).
func productExceptSelf(nums []int) []int {
	n := len(nums)
	out := make([]int, n)

	out[0] = 1
	for i := 1; i < n; i++ {
		out[i] = out[i-1] * nums[i-1]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		out[i] *= suffix
		suffix *= nums[i]
	}
	return out
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))        // [24 12 8 6]
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))    // [0 0 9 0 0]
}
