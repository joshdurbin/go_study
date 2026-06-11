//go:build ignore

package main

import "fmt"

// wordBreak reports whether s can be segmented into a sequence of dictionary words.
// Words may be reused. Target O(n^2) time with 1-D DP.
func wordBreak(s string, dict []string) bool {
	// TODO: implement
	return false
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))                      // expect true
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))                 // expect true
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})) // expect false
}
