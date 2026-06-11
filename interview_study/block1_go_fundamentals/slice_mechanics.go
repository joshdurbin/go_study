//go:build ignore

package main

import "fmt"

// removeDuplicates removes duplicates from a slice in-place, preserving order.
// Uses a write-head pointer and a seen map. O(n) time, O(n) space.
func removeDuplicates(a []int) []int {
	seen := make(map[int]bool, len(a))
	write := 0
	for _, v := range a {
		if !seen[v] {
			seen[v] = true
			a[write] = v
			write++
		}
	}
	return a[:write]
}

func main() {
	input := []int{1, 2, 2, 3, 1, 4}
	fmt.Println(removeDuplicates(input)) // [1 2 3 4]

	input2 := []int{5, 5, 5}
	fmt.Println(removeDuplicates(input2)) // [5]
}
