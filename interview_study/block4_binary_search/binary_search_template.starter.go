//go:build ignore

package main

import "fmt"

// search returns the index of target in a sorted slice, or -1 if absent.
// Use closed-interval binary search: lo, hi := 0, len(a)-1.
func search(a []int, target int) int {
	// TODO: implement
	return -1
}

// firstTrue returns the smallest index in [0, n) for which pred(i) is true,
// or n if no such index exists. Assumes pred is monotone (false...false true...true).
func firstTrue(n int, pred func(int) bool) int {
	// TODO: implement (boundary-search variant)
	return n
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))  // expect 4
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))  // expect -1
	fmt.Println(firstTrue(10, func(i int) bool { return i*i >= 25 })) // expect 5
}
