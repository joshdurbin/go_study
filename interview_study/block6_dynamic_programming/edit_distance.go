//go:build ignore

package main

import "fmt"

// editDistance: 2-D DP. dp[i][j] = edit distance between a[:i] and b[:j].
// O(m*n) time and space.
func editDistance(a, b string) int {
	m, n := len(a), len(b)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i // a[:i] → "" needs i deletions
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // "" → b[:j] needs j insertions
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min3(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}
	return dp[m][n]
}

func min3(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}

func main() {
	fmt.Println(editDistance("horse", "ros"))          // 3
	fmt.Println(editDistance("intention", "execution")) // 5
}
