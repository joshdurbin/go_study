//go:build ignore

package main

import "fmt"

// rotate by k in-place: reverse all, reverse [0..k), reverse [k..n). O(n), O(1).
func rotate(a []int, k int) {
	n := len(a)
	if n == 0 {
		return
	}
	k %= n
	reverse(a, 0, n-1)
	reverse(a, 0, k-1)
	reverse(a, k, n-1)
}

func reverse(a []int, i, j int) {
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(a, 3)
	fmt.Println(a) // [5 6 7 1 2 3 4]

	b := []int{-1, -100, 3, 99}
	rotate(b, 2)
	fmt.Println(b) // [3 99 -1 -100]
}
