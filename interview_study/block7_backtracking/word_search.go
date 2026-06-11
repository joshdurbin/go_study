//go:build ignore

package main

import "fmt"

// exist: DFS from every cell. Mark visited by overwriting the board cell with
// '#' (sentinel) and restore it on backtrack — avoids allocating a separate
// visited[][]. O(m*n*4^L) time, O(L) recursion depth where L = len(word).
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	var dfs func(r, c, k int) bool
	dfs = func(r, c, k int) bool {
		if k == len(word) {
			return true
		}
		if r < 0 || r >= m || c < 0 || c >= n || board[r][c] != word[k] {
			return false
		}
		saved := board[r][c]
		board[r][c] = '#'
		found := dfs(r+1, c, k+1) || dfs(r-1, c, k+1) ||
			dfs(r, c+1, k+1) || dfs(r, c-1, k+1)
		board[r][c] = saved
		return found
	}
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	b1 := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(b1, "ABCCED")) // true
	fmt.Println(exist(b1, "SEE"))    // true
	fmt.Println(exist(b1, "ABCB"))   // false
}
