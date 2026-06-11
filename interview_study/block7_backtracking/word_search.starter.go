//go:build ignore

package main

import "fmt"

// exist returns true if word can be formed by a path of orthogonally adjacent
// cells in board, with no cell reused. Try the "mark cell with '#' then
// restore" trick to avoid allocating a visited[][].
func exist(board [][]byte, word string) bool {
	// TODO: implement
	return false
}

func main() {
	b1 := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(b1, "ABCCED")) // expect true
	fmt.Println(exist(b1, "SEE"))    // expect true
	fmt.Println(exist(b1, "ABCB"))   // expect false
}
