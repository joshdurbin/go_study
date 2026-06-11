//go:build ignore

package main

import "fmt"

// largestRectangleArea: monotonic increasing stack of indices. When the next
// bar is shorter than the stack top, the top bar can extend no further right,
// so pop it and compute its rectangle: height * (i - newTop - 1). The "newTop"
// after popping is the first shorter bar on the left, giving the exact width.
// A sentinel 0 appended forces all remaining bars to drain.
// O(n) time, O(n) space.
func largestRectangleArea(heights []int) int {
	stack := make([]int, 0, len(heights)+1) // increasing heights, indices
	best := 0
	// iterate one extra step with height 0 to flush
	for i := 0; i <= len(heights); i++ {
		var h int
		if i == len(heights) {
			h = 0
		} else {
			h = heights[i]
		}
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			if area := heights[top] * width; area > best {
				best = area
			}
		}
		stack = append(stack, i)
	}
	return best
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3})) // 10
	fmt.Println(largestRectangleArea([]int{2, 4}))             // 4
	fmt.Println(largestRectangleArea([]int{1, 1, 1, 1}))       // 4
	fmt.Println(largestRectangleArea([]int{6, 7, 5, 2, 4, 5, 9, 3})) // 16
}
