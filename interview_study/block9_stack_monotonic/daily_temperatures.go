//go:build ignore

package main

import "fmt"

// dailyTemperatures: monotonic decreasing stack of indices. When the current
// temp breaks the decreasing chain, pop everything it beats and record the
// distance. Each index is pushed and popped at most once.
// O(n) time, O(n) space.
func dailyTemperatures(temps []int) []int {
	result := make([]int, len(temps))
	stack := make([]int, 0, len(temps)) // indices, temps[stack[k]] decreasing
	for i, t := range temps {
		for len(stack) > 0 && temps[stack[len(stack)-1]] < t {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[top] = i - top
		}
		stack = append(stack, i)
	}
	return result
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})) // [1 1 4 2 1 1 0 0]
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))                 // [1 1 1 0]
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))                     // [1 1 0]
}
