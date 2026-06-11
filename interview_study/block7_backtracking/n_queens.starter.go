//go:build ignore

package main

import "fmt"

// solveNQueens returns every distinct placement of n queens on an n×n board
// so that none attack each other. Each solution is an []string where 'Q'
// marks a queen and '.' an empty square. Iterate by row and track three
// attack sets (columns, '\' diagonals, '/' diagonals) for O(1) checks.
func solveNQueens(n int) [][]string {
	// TODO: implement
	return nil
}

func main() {
	fmt.Println(solveNQueens(4)) // expect [[.Q.. ...Q Q... ..Q.] [..Q. Q... ...Q .Q..]]
	fmt.Println(solveNQueens(1)) // expect [[Q]]
}
