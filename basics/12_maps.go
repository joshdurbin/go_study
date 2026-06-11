//go:build ignore

package main

import "fmt"

func main() {
	// Initialization — always initialize before writing (nil map panics on write)
	m := make(map[string]int)
	m["alpha"] = 1
	m["beta"] = 2
	m["gamma"] = 3

	// Map literal
	scores := map[string]int{
		"alice": 95,
		"bob":   82,
		"carol": 91,
	}

	// Read with comma-ok to check existence
	v, ok := scores["alice"]
	fmt.Println(v, ok) // 95 true

	v, ok = scores["dave"]
	fmt.Println(v, ok) // 0 false — zero value, not present

	// Delete a key
	delete(scores, "bob")
	fmt.Println(len(scores)) // 2

	// Iterate (order is random — never rely on it)
	for k, v := range m {
		fmt.Printf("%s: %d\n", k, v)
	}

	// Map as a set
	seen := make(map[string]struct{})
	words := []string{"go", "is", "go", "great"}
	for _, w := range words {
		seen[w] = struct{}{}
	}
	fmt.Println(len(seen)) // 3 unique words

	// Nested maps
	graph := map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
	}
	fmt.Println(graph["a"]) // [b c]

	// Counting frequency (classic pattern)
	freq := make(map[rune]int)
	for _, c := range "mississippi" {
		freq[c]++
	}
	fmt.Println(freq) // map[i:4 m:1 p:2 s:4]
}
