//go:build ignore

package main

import "fmt"

// trap: two-pointer sweep. Whichever side has the lower running max is the
// binding wall for THAT index, so the answer there is leftMax/rightMax - h.
// O(n) time, O(1) space.
func trap(h []int) int {
	left, right := 0, len(h)-1
	lMax, rMax, total := 0, 0, 0
	for left < right {
		if h[left] < h[right] {
			if h[left] >= lMax {
				lMax = h[left]
			} else {
				total += lMax - h[left]
			}
			left++
		} else {
			if h[right] >= rMax {
				rMax = h[right]
			} else {
				total += rMax - h[right]
			}
			right--
		}
	}
	return total
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))                   // 9
}
