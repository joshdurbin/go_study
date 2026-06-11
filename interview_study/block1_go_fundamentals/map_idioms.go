//go:build ignore

package main

import (
	"fmt"
	"strings"
)

// wordFrequency returns a frequency map for space-separated words.
func wordFrequency(s string) map[string]int {
	freq := make(map[string]int)
	for _, w := range strings.Fields(s) {
		freq[w]++ // safe: zero value of int is 0
	}
	return freq
}

// topWord returns the most frequent word; ties broken lexicographically.
func topWord(freq map[string]int) string {
	best := ""
	for word, count := range freq {
		if best == "" {
			best = word
			continue
		}
		if count > freq[best] || (count == freq[best] && word < best) {
			best = word
		}
	}
	return best
}

func main() {
	s := "the cat sat on the mat the cat"
	freq := wordFrequency(s)
	fmt.Println(freq)          // map[cat:2 mat:1 on:1 sat:1 the:3]
	fmt.Println(topWord(freq)) // the
}
