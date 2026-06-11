//go:build ignore

package main

import "fmt"

// wordBreak: 1-D DP. dp[i] = can we segment s[:i]?
// For each i, scan j<i; if dp[j] and s[j:i] in dict, dp[i] = true.
// O(n²) time × substring/map lookup, O(n) space.
func wordBreak(s string, dict []string) bool {
	set := make(map[string]struct{}, len(dict))
	for _, w := range dict {
		set[w] = struct{}{}
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] {
				if _, ok := set[s[j:i]]; ok {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[len(s)]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))                   // true
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))              // true
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})) // false
}
