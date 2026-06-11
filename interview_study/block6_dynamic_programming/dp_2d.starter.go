//go:build ignore

package main

import "fmt"

// uniquePaths returns the number of distinct paths from the top-left to the
// bottom-right of an m×n grid, moving only right or down.
// 2-D DP: dp[i][j] = dp[i-1][j] + dp[i][j-1].
func uniquePaths(m, n int) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println(uniquePaths(3, 7)) // expect 28
	fmt.Println(uniquePaths(3, 2)) // expect 3
	fmt.Println(uniquePaths(1, 1)) // expect 1
}
