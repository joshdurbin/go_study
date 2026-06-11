//go:build ignore

package main

import "fmt"

// coinChange: bottom-up DP. Sentinel value = amount+1 acts as "infinity".
// O(amount * len(coins)) time, O(amount) space.
func coinChange(coins []int, amount int) int {
	const INF = int(1e9)
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = INF
		for _, c := range coins {
			if c <= i && dp[i-c]+1 < dp[i] {
				dp[i] = dp[i-c] + 1
			}
		}
	}
	if dp[amount] == INF {
		return -1
	}
	return dp[amount]
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11)) // 3
	fmt.Println(coinChange([]int{2}, 3))        // -1
	fmt.Println(coinChange([]int{1}, 0))        // 0
}
