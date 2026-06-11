//go:build ignore

package main

import "fmt"

func maxWater(heights []int) int {
	lo, hi := 0, len(heights)-1
	maxArea := 0
	for lo < hi {
		h := heights[lo]
		if heights[hi] < h {
			h = heights[hi]
		}
		area := h * (hi - lo)
		if area > maxArea {
			maxArea = area
		}
		if heights[lo] < heights[hi] {
			lo++
		} else {
			hi--
		}
	}
	return maxArea
}

func main() {
	fmt.Println(maxWater([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	fmt.Println(maxWater([]int{1, 1}))                        // 1
}
