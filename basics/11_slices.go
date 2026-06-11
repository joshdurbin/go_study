//go:build ignore

package main

import "fmt"

func main() {
	// Declaration and initialization
	var s []int               // nil slice (len=0, cap=0)
	s2 := []int{1, 2, 3}     // slice literal
	s3 := make([]int, 3, 5)  // len=3, cap=5

	fmt.Println(s == nil, len(s))        // true 0
	fmt.Println(s2, len(s2), cap(s2))   // [1 2 3] 3 3
	fmt.Println(s3, len(s3), cap(s3))   // [0 0 0] 3 5

	// append may allocate a new backing array when capacity is exceeded
	s3 = append(s3, 4, 5, 6)
	fmt.Println(s3, len(s3), cap(s3)) // [0 0 0 4 5 6] 6 10 (cap doubled)

	// Slicing a slice: shares the backing array!
	a := []int{0, 1, 2, 3, 4}
	b := a[1:4] // [1 2 3] — shared memory with a
	b[0] = 99
	fmt.Println(a) // [0 99 2 3 4] — a was mutated!

	// Copy to avoid shared backing array
	c := make([]int, len(b))
	copy(c, b)
	c[0] = 0
	fmt.Println(b) // [99 2 3] — b unchanged

	// 2D slice
	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 3)
		matrix[i][i] = 1 // identity matrix
	}
	fmt.Println(matrix) // [[1 0 0] [0 1 0] [0 0 1]]

	// Delete element at index i (order-preserving)
	remove := func(s []int, i int) []int {
		return append(s[:i], s[i+1:]...)
	}
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(remove(nums, 2)) // [1 2 4 5]
}
