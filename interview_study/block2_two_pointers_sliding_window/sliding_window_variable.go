//go:build ignore

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)
	best := 0
	left := 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		if idx, ok := charIndex[c]; ok && idx >= left {
			left = idx + 1
		}
		charIndex[c] = right
		if w := right - left + 1; w > best {
			best = w
		}
	}
	return best
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
}
