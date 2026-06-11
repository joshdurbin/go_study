//go:build ignore

package main

import "fmt"

func subarraySum(nums []int, k int) int {
	prefixCount := map[int]int{0: 1}
	sum := 0
	count := 0
	for _, n := range nums {
		sum += n
		count += prefixCount[sum-k]
		prefixCount[sum]++
	}
	return count
}

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))   // 2
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))   // 2
	fmt.Println(subarraySum([]int{-1, -1, 1}, 0)) // 1
}
