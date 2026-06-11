//go:build ignore

package main

import "fmt"

// isPowerOfTwo: a power of two has exactly one bit set. n-1 flips that bit and
// turns every lower zero into one, so n & (n-1) == 0 iff n was a single-bit value.
// Guard n > 0 because 0 also satisfies n & (n-1) == 0 trivially.
// O(1) time, O(1) space.
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func main() {
	fmt.Println(isPowerOfTwo(1))  // true   (2^0)
	fmt.Println(isPowerOfTwo(16)) // true   (2^4)
	fmt.Println(isPowerOfTwo(3))  // false
	fmt.Println(isPowerOfTwo(0))  // false
	fmt.Println(isPowerOfTwo(-4)) // false
}
