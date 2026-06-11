//go:build ignore

package main

import "fmt"

// lengthOfLongestSubstring returns the length of the longest substring of s
// containing no repeated characters. Target O(n) time with a sliding window.
func lengthOfLongestSubstring(s string) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // expect 3 ("abc")
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // expect 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // expect 3 ("wke")
	fmt.Println(lengthOfLongestSubstring(""))         // expect 0
}
