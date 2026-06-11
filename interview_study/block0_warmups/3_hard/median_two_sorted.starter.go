//go:build ignore

package main

import "fmt"

// findMedianSortedArrays returns the median of the merged sorted set of a and b.
// Target O(log(min(m, n))) time — binary-search the partition position in the
// shorter array.
func findMedianSortedArrays(a, b []int) float64 {
	// TODO: implement
	return 0
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))           // expect 2
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))        // expect 2.5
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))               // expect 1
	fmt.Println(findMedianSortedArrays([]int{1, 3, 8}, []int{7, 9, 10, 11})) // expect 8
}
