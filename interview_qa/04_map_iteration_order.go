//go:build ignore

package main

import (
	"fmt"
	"sort"
)

// Map iteration order is intentionally randomized by the Go runtime to discourage
// callers from depending on it. If you need determinism, sort the keys.

func main() {
	m := map[string]int{
		"alpha":   1,
		"bravo":   2,
		"charlie": 3,
		"delta":   4,
		"echo":    5,
	}

	// ─── Case 1: two ranges, two (likely) different orders ─────
	fmt.Print("range 1: ")
	for k := range m {
		fmt.Print(k, " ")
	}
	fmt.Println()

	fmt.Print("range 2: ")
	for k := range m {
		fmt.Print(k, " ")
	}
	fmt.Println()
	// Same program run, same map — different orders. Re-run the program to see
	// the seed change too.

	// ─── Case 2: deterministic — sort keys first ─────
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Print("sorted:  ")
	for _, k := range keys {
		fmt.Printf("%s=%d ", k, m[k])
	}
	fmt.Println()
}
