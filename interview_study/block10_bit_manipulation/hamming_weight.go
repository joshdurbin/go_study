//go:build ignore

package main

import "fmt"

// hammingWeight: Kernighan's trick. n & (n-1) clears the lowest set bit each
// iteration, so the loop runs exactly k times for k bits set — not 32.
// O(k) time where k = number of set bits, O(1) space.
func hammingWeight(n uint32) int {
	count := 0
	for n != 0 {
		n &= n - 1
		count++
	}
	return count
}

func main() {
	fmt.Println(hammingWeight(0b00000000000000000000000000001011)) // 3
	fmt.Println(hammingWeight(0b00000000000000000000000010000000)) // 1
	fmt.Println(hammingWeight(0b11111111111111111111111111111101)) // 31
}
