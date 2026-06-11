//go:build ignore

package main

import "fmt"

// combinations returns all k-element combinations of [1..n] in ascending order.
// Use a backtracking template with a start index; prune the loop bound so you
// don't recurse into branches that cannot fill the path.
func combinations(n, k int) [][]int {
	// TODO: implement
	return nil
}

func main() {
	fmt.Println(combinations(4, 2)) // expect [[1 2] [1 3] [1 4] [2 3] [2 4] [3 4]]
	fmt.Println(combinations(3, 1)) // expect [[1] [2] [3]]
}
