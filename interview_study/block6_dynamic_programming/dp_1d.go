//go:build ignore

package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // sentinel: impossible
	}
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i && dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// climbStairs: number of distinct ways to reach step n (1 or 2 steps at a time).
// Space-optimized: only track last two values.
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	prev2, prev1 := 1, 2
	for i := 3; i <= n; i++ {
		prev2, prev1 = prev1, prev1+prev2
	}
	return prev1
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11)) // 3
	fmt.Println(coinChange([]int{2}, 3))         // -1
	fmt.Println(coinChange([]int{1}, 0))         // 0

	fmt.Println(climbStairs(2)) // 2
	fmt.Println(climbStairs(5)) // 8
}
