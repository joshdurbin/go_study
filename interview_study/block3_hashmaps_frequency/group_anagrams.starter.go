//go:build ignore

package main

import "fmt"

// groupAnagrams groups words that are anagrams of each other.
// Output: a slice of groups, each containing the original words.
func groupAnagrams(words []string) [][]string {
	// TODO: implement — pick a canonical key per word and group by it.
	return nil
}

func main() {
	got := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(got)
	// expect (in some order): [[eat tea ate] [tan nat] [bat]]
}
