//go:build ignore

package main

import "fmt"

// rotate rotates a to the right by k steps, in-place, using O(1) extra space.
func rotate(a []int, k int) {
	// TODO: implement
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(a, 3)
	fmt.Println(a) // expect [5 6 7 1 2 3 4]

	b := []int{-1, -100, 3, 99}
	rotate(b, 2)
	fmt.Println(b) // expect [3 99 -1 -100]
}
