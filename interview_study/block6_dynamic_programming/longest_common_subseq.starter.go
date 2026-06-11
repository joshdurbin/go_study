//go:build ignore

package main

import "fmt"

// lcs returns the length of the Longest Common Subsequence of a and b.
// (A subsequence keeps order but does not need to be contiguous.)
// 2-D DP: dp[i][j] over prefixes a[:i] and b[:j].
func lcs(a, b string) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println(lcs("abcde", "ace")) // expect 3
	fmt.Println(lcs("abc", "abc"))   // expect 3
	fmt.Println(lcs("abc", "def"))   // expect 0
}
