//go:build ignore

package main

import "fmt"

// combinations: backtrack with a start index. Record when path length == k.
// Pruning: stop iterating when remaining elements can't fill the path.
// O(k * C(n,k)) time, O(k) recursion depth.
func combinations(n, k int) [][]int {
	res := [][]int{}
	path := []int{}
	var backtrack func(start int)
	backtrack = func(start int) {
		if len(path) == k {
			cp := make([]int, k)
			copy(cp, path)
			res = append(res, cp)
			return
		}
		need := k - len(path)
		for i := start; i <= n-need+1; i++ {
			path = append(path, i)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(1)
	return res
}

func main() {
	fmt.Println(combinations(4, 2)) // [[1 2] [1 3] [1 4] [2 3] [2 4] [3 4]]
	fmt.Println(combinations(3, 1)) // [[1] [2] [3]]
}
