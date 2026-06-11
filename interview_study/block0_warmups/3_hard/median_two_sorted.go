//go:build ignore

package main

import (
	"fmt"
	"math"
)

// findMedianSortedArrays: binary-search the partition in the SHORTER array.
// For partition i in A and j = half - i in B, valid when A[i-1] <= B[j] and B[j-1] <= A[i].
// O(log(min(m,n))) time, O(1) space.
func findMedianSortedArrays(a, b []int) float64 {
	if len(a) > len(b) {
		a, b = b, a
	}
	m, n := len(a), len(b)
	half := (m + n + 1) / 2

	lo, hi := 0, m
	for lo <= hi {
		i := (lo + hi) / 2
		j := half - i

		aLeft := math.MinInt
		if i > 0 {
			aLeft = a[i-1]
		}
		aRight := math.MaxInt
		if i < m {
			aRight = a[i]
		}
		bLeft := math.MinInt
		if j > 0 {
			bLeft = b[j-1]
		}
		bRight := math.MaxInt
		if j < n {
			bRight = b[j]
		}

		if aLeft <= bRight && bLeft <= aRight {
			if (m+n)%2 == 1 {
				return float64(max(aLeft, bLeft))
			}
			return float64(max(aLeft, bLeft)+min(aRight, bRight)) / 2.0
		} else if aLeft > bRight {
			hi = i - 1
		} else {
			lo = i + 1
		}
	}
	return 0
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))     // 2
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))  // 2.5
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))         // 1
	fmt.Println(findMedianSortedArrays([]int{1, 3, 8}, []int{7, 9, 10, 11})) // 8
}
