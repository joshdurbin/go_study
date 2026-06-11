//go:build ignore

package main

import "fmt"

// largestRectangleArea returns the area of the largest rectangle that fits
// inside the histogram given by heights (each bar has width 1).
// Hint: monotonic increasing stack of indices; on pop, width = i - newTop - 1.
// Append a sentinel 0 at the end to drain the stack.
func largestRectangleArea(heights []int) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))       // expect 10
	fmt.Println(largestRectangleArea([]int{2, 4}))                   // expect 4
	fmt.Println(largestRectangleArea([]int{1, 1, 1, 1}))             // expect 4
	fmt.Println(largestRectangleArea([]int{6, 7, 5, 2, 4, 5, 9, 3})) // expect 16
}
