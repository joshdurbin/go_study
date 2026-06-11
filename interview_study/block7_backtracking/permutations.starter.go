//go:build ignore

package main

import "fmt"

// permutations returns every ordering of the distinct ints in nums. There are
// n! results. Try the swap-in-place approach — it avoids allocating a used[]
// set per call frame.
func permutations(nums []int) [][]int {
	// TODO: implement
	return nil
}

func main() {
	fmt.Println(permutations([]int{1, 2, 3})) // expect [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 2 1] [3 1 2]]
	fmt.Println(permutations([]int{0, 1}))    // expect [[0 1] [1 0]]
}
