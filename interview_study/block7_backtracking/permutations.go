//go:build ignore

package main

import "fmt"

// permutations: swap-in-place backtracking. For position `first`, swap each
// nums[i] (i >= first) into place, recurse, then swap back.
// O(n * n!) time, O(n) recursion depth.
func permutations(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	var backtrack func(first int)
	backtrack = func(first int) {
		if first == n {
			cp := make([]int, n)
			copy(cp, nums)
			res = append(res, cp)
			return
		}
		for i := first; i < n; i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backtrack(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
	backtrack(0)
	return res
}

func main() {
	fmt.Println(permutations([]int{1, 2, 3})) // [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 2 1] [3 1 2]]
	fmt.Println(permutations([]int{0, 1}))    // [[0 1] [1 0]]
}
