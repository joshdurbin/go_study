//go:build ignore

package main

import "fmt"

// spiralOrder returns all elements of m in clockwise spiral order starting at
// the top-left corner.
func spiralOrder(m [][]int) []int {
	// TODO: implement
	return nil
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	// expect [1 2 3 6 9 8 7 4 5]
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	// expect [1 2 3 4 8 12 11 10 9 5 6 7]
}
