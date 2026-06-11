//go:build ignore

package main

import "fmt"

// knapsack01: 1-D rolling array. Iterating weight RIGHT TO LEFT prevents using
// the same item twice (each cell still reflects the previous item's row).
// O(n * W) time, O(W) space.
func knapsack01(weights, values []int, capacity int) int {
	dp := make([]int, capacity+1)
	for i, w := range weights {
		for c := capacity; c >= w; c-- {
			if dp[c-w]+values[i] > dp[c] {
				dp[c] = dp[c-w] + values[i]
			}
		}
	}
	return dp[capacity]
}

func main() {
	fmt.Println(knapsack01([]int{1, 3, 4, 5}, []int{1, 4, 5, 7}, 7)) // 9
	fmt.Println(knapsack01([]int{2, 3, 4, 5}, []int{3, 4, 5, 6}, 5)) // 7
}
