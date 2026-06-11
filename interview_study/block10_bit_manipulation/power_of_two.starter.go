//go:build ignore

package main

import "fmt"

// isPowerOfTwo returns true iff n is a positive power of two.
// Target O(1) — no loops, no recursion.
func isPowerOfTwo(n int) bool {
	// TODO: implement
	return false
}

func main() {
	fmt.Println(isPowerOfTwo(1))  // expect true
	fmt.Println(isPowerOfTwo(16)) // expect true
	fmt.Println(isPowerOfTwo(3))  // expect false
	fmt.Println(isPowerOfTwo(0))  // expect false
	fmt.Println(isPowerOfTwo(-4)) // expect false
}
