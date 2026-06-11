//go:build ignore

package main

import "fmt"

// subsets: classic include/exclude backtracking. At each index, choose to take
// nums[i] or skip it. Append a copy of the current path at every node.
// O(n * 2^n) time, O(n) extra space for recursion.
func subsets(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	var backtrack func(start int)
	backtrack = func(start int) {
		cp := make([]int, len(path))
		copy(cp, path)
		res = append(res, cp)
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3})) // [[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]
	fmt.Println(subsets([]int{0}))       // [[] [0]]
}
