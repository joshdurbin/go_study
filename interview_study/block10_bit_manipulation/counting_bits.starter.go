//go:build ignore

package main

import "fmt"

// countBits returns a slice of length n+1 where result[i] is the popcount of i.
// Target O(n) total work.
func countBits(n int) []int {
	// TODO: implement
	return nil
}

func main() {
	fmt.Println(countBits(2)) // expect [0 1 1]
	fmt.Println(countBits(5)) // expect [0 1 1 2 1 2]
	fmt.Println(countBits(0)) // expect [0]
}
