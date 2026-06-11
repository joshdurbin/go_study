//go:build ignore

package main

import "fmt"

// spiralOrder: shrink boundaries inward as each edge is consumed.
// Guard rows/cols left after a horizontal/vertical pass to avoid re-visiting
// the same line in a single-row or single-column remainder.
func spiralOrder(m [][]int) []int {
	if len(m) == 0 {
		return nil
	}
	top, bottom := 0, len(m)-1
	left, right := 0, len(m[0])-1
	out := make([]int, 0, len(m)*len(m[0]))
	for top <= bottom && left <= right {
		for c := left; c <= right; c++ { // top row L→R
			out = append(out, m[top][c])
		}
		top++
		for r := top; r <= bottom; r++ { // right col T→B
			out = append(out, m[r][right])
		}
		right--
		if top <= bottom {
			for c := right; c >= left; c-- { // bottom row R→L
				out = append(out, m[bottom][c])
			}
			bottom--
		}
		if left <= right {
			for r := bottom; r >= top; r-- { // left col B→T
				out = append(out, m[r][left])
			}
			left++
		}
	}
	return out
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}
