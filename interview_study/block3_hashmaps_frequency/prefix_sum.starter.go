//go:build ignore

package main

import "fmt"

// subarraySumK returns the number of contiguous subarrays of a summing to k.
// Target O(n) time using prefix sums + a hashmap of counts.
func subarraySumK(a []int, k int) int {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println(subarraySumK([]int{1, 1, 1}, 2))     // expect 2
	fmt.Println(subarraySumK([]int{1, 2, 3}, 3))     // expect 2
	fmt.Println(subarraySumK([]int{-1, -1, 1}, 0))   // expect 1
}
