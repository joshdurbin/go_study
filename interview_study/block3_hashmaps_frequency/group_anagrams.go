//go:build ignore

package main

import (
	"fmt"
	"sort"
)

// groupAnagramsSort uses sorted string as map key.
func groupAnagramsSort(words []string) [][]string {
	groups := make(map[string][]string)
	for _, w := range words {
		key := sortString(w)
		groups[key] = append(groups[key], w)
	}
	result := make([][]string, 0, len(groups))
	for _, g := range groups {
		result = append(result, g)
	}
	return result
}

func sortString(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}

// groupAnagramsFreq uses [26]int frequency array as map key. O(n*k).
func groupAnagramsFreq(words []string) [][]string {
	groups := make(map[[26]int][]string)
	for _, w := range words {
		var key [26]int
		for _, c := range w {
			key[c-'a']++
		}
		groups[key] = append(groups[key], w)
	}
	result := make([][]string, 0, len(groups))
	for _, g := range groups {
		result = append(result, g)
	}
	return result
}

func main() {
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagramsSort(words))
	fmt.Println(groupAnagramsFreq(words))
}
