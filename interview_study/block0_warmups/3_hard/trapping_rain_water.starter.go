//go:build ignore

package main

import "fmt"

// trap returns the total water trapped above the bar heights h after rain.
// Each bar has width 1. Target O(n) time, O(1) space with two pointers.
func trap(h []int) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // expect 6
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))                   // expect 9
}
