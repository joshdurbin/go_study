//go:build ignore

package main

import (
	"fmt"
	"strings"
)

// solveNQueens: place one queen per row. Track three attack sets — columns,
// `\` diagonals (r-c, offset by n-1 to be non-negative), and `/` diagonals
// (r+c). Constant-time conflict check beats scanning prior rows.
// O(n!) time (branching factor shrinks), O(n) recursion + O(n) sets.
func solveNQueens(n int) [][]string {
	res := [][]string{}
	cols := make([]bool, n)
	diag := make([]bool, 2*n-1) // r - c + (n-1)
	anti := make([]bool, 2*n-1) // r + c
	queens := make([]int, n)    // queens[r] = column of queen in row r

	var backtrack func(r int)
	backtrack = func(r int) {
		if r == n {
			board := make([]string, n)
			for i := 0; i < n; i++ {
				row := strings.Repeat(".", queens[i]) + "Q" + strings.Repeat(".", n-queens[i]-1)
				board[i] = row
			}
			res = append(res, board)
			return
		}
		for c := 0; c < n; c++ {
			d := r - c + (n - 1)
			a := r + c
			if cols[c] || diag[d] || anti[a] {
				continue
			}
			queens[r] = c
			cols[c], diag[d], anti[a] = true, true, true
			backtrack(r + 1)
			cols[c], diag[d], anti[a] = false, false, false
		}
	}
	backtrack(0)
	return res
}

func main() {
	fmt.Println(solveNQueens(4)) // [[.Q.. ...Q Q... ..Q.] [..Q. Q... ...Q .Q..]]
	fmt.Println(solveNQueens(1)) // [[Q]]
}
