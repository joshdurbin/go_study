//go:build ignore

package main

import "fmt"

// canJump: track the farthest index reachable so far; if i ever exceeds it, we're stuck.
// O(n) time, O(1) space.
func canJump(nums []int) bool {
	farthest := 0
	for i, x := range nums {
		if i > farthest {
			return false
		}
		if i+x > farthest {
			farthest = i + x
		}
		if farthest >= len(nums)-1 {
			return true
		}
	}
	return true
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4})) // true
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) // false
	fmt.Println(canJump([]int{0}))             // true
	fmt.Println(canJump([]int{2, 0, 0}))       // true
}
