//go:build ignore

package main

import "fmt"

// coinChange returns the minimum number of coins needed to make `amount`,
// or -1 if it cannot be made. Coins can be used unlimited times.
// 1-D DP: dp[i] = min over coins c<=i of (dp[i-c] + 1).
func coinChange(coins []int, amount int) int {
	// TODO: implement
	return -1
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11)) // expect 3 (5+5+1)
	fmt.Println(coinChange([]int{2}, 3))        // expect -1
	fmt.Println(coinChange([]int{1}, 0))        // expect 0
}
