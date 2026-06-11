//go:build ignore

package main

import "fmt"

// topKFrequent returns the k most frequent integers in nums, in any order.
// Target O(n log k) using a min-heap, or O(n) using bucket sort.
func topKFrequent(nums []int, k int) []int {
	// TODO: implement
	return nil
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)) // expect [1 2] (any order)
	fmt.Println(topKFrequent([]int{1}, 1))                // expect [1]
}
