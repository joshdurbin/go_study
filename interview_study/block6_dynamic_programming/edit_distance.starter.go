//go:build ignore

package main

import "fmt"

// editDistance returns the minimum number of insertions, deletions, or
// substitutions to convert a into b.
// 2-D DP. Don't forget the base cases for prefix→empty and empty→prefix.
func editDistance(a, b string) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println(editDistance("horse", "ros"))           // expect 3
	fmt.Println(editDistance("intention", "execution")) // expect 5
	fmt.Println(editDistance("", "abc"))                 // expect 3
	fmt.Println(editDistance("abc", ""))                 // expect 3
}
