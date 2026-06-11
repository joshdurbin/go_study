//go:build ignore

package main

import "fmt"

// shipWithinDays returns the minimum ship capacity that allows all weights
// (in order) to be shipped within D days. Each day's load must not exceed capacity.
// Binary search the ANSWER (capacity), not an index.
func shipWithinDays(weights []int, D int) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)) // expect 15
	fmt.Println(shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))              // expect 6
	fmt.Println(shipWithinDays([]int{1, 2, 3, 1, 1}, 4))                 // expect 3
}
