//go:build ignore

package main

import "fmt"

// hammingWeight returns the number of 1-bits in n.
// Target O(k) where k is the number of set bits.
func hammingWeight(n uint32) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println(hammingWeight(0b00000000000000000000000000001011)) // expect 3
	fmt.Println(hammingWeight(0b00000000000000000000000010000000)) // expect 1
	fmt.Println(hammingWeight(0b11111111111111111111111111111101)) // expect 31
}
