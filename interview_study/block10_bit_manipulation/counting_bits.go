//go:build ignore

package main

import "fmt"

// countBits: DP using a prior result. For any i > 0, dropping the lowest bit
// gives i>>1 — which we've already counted. So popcount(i) = popcount(i>>1) + (i&1).
// O(n) time, O(n) space for the output.
func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 1; i <= n; i++ {
		result[i] = result[i>>1] + (i & 1)
	}
	return result
}

func main() {
	fmt.Println(countBits(2)) // [0 1 1]
	fmt.Println(countBits(5)) // [0 1 1 2 1 2]
	fmt.Println(countBits(0)) // [0]
}
