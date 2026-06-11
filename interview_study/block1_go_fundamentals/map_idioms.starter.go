//go:build ignore

package main

import "fmt"

// wordCount returns a map of word → occurrence count.
// Demonstrate: zero-value-friendly increments, comma-ok presence checks.
func wordCount(words []string) map[string]int {
	// TODO: implement
	return nil
}

// firstUnique returns the first word with count 1, or "" if none.
// Demonstrates ordered iteration over a non-ordered structure.
func firstUnique(words []string) string {
	// TODO: implement
	return ""
}

func main() {
	counts := wordCount([]string{"go", "is", "go", "go", "fast"})
	fmt.Println(counts) // expect map[fast:1 go:3 is:1]
	fmt.Println(firstUnique([]string{"go", "is", "go", "fast", "is"})) // expect "fast"
}
