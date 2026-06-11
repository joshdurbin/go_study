//go:build ignore

package main

import "fmt"

// nextGreaterElements: circular array → walk indices 0..2n-1 and index with
// `i % n`. Monotonic decreasing stack of indices; when the current value beats
// the stack top, that popped index has found its next greater.
// O(n) time, O(n) space.
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i := range result {
		result[i] = -1
	}
	stack := make([]int, 0, n) // indices, nums[stack[k]] decreasing
	for i := 0; i < 2*n; i++ {
		v := nums[i%n]
		for len(stack) > 0 && nums[stack[len(stack)-1]] < v {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[top] = v
		}
		if i < n {
			stack = append(stack, i) // only push originals; second pass just resolves
		}
	}
	return result
}

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))       // [2 -1 2]
	fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 3})) // [2 3 4 -1 4]
	fmt.Println(nextGreaterElements([]int{5, 4, 3, 2, 1})) // [-1 5 5 5 5]
}
