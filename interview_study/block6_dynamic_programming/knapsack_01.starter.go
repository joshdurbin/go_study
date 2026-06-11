//go:build ignore

package main

import "fmt"

// knapsack01 returns the maximum total value such that the total weight ≤ capacity.
// Each item may be taken at most once.
// 1-D rolling array trick: iterate weight RIGHT-TO-LEFT in the inner loop.
func knapsack01(weights, values []int, capacity int) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println(knapsack01([]int{1, 3, 4, 5}, []int{1, 4, 5, 7}, 7)) // expect 9
	fmt.Println(knapsack01([]int{2, 3, 4, 5}, []int{3, 4, 5, 6}, 5)) // expect 7
}
