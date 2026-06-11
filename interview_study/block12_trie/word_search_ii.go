//go:build ignore

package main

import (
	"fmt"
	"sort"
)

// wordSearchII: build a trie of the dictionary, DFS the board carrying a
// pointer into the trie. Trie pruning beats running Word Search I per word.
// O(m*n * 4^L) worst case where L is the longest word; in practice far less
// because failed prefixes prune entire subtrees.

type tnode struct {
	children [26]*tnode
	word     string // non-empty marks a complete word terminating here
}

func buildTrie(words []string) *tnode {
	root := &tnode{}
	for _, w := range words {
		n := root
		for i := 0; i < len(w); i++ {
			c := w[i] - 'a'
			if n.children[c] == nil {
				n.children[c] = &tnode{}
			}
			n = n.children[c]
		}
		n.word = w
	}
	return root
}

func findWords(board [][]byte, words []string) []string {
	root := buildTrie(words)
	var out []string
	m, n := len(board), len(board[0])
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			dfs(board, r, c, root, &out)
		}
	}
	sort.Strings(out)
	return out
}

func dfs(board [][]byte, r, c int, n *tnode, out *[]string) {
	ch := board[r][c]
	if ch == '#' || n.children[ch-'a'] == nil {
		return
	}
	n = n.children[ch-'a']
	if n.word != "" {
		*out = append(*out, n.word)
		n.word = "" // dedupe — never report the same word twice
	}
	board[r][c] = '#'
	if r > 0 {
		dfs(board, r-1, c, n, out)
	}
	if r+1 < len(board) {
		dfs(board, r+1, c, n, out)
	}
	if c > 0 {
		dfs(board, r, c-1, n, out)
	}
	if c+1 < len(board[0]) {
		dfs(board, r, c+1, n, out)
	}
	board[r][c] = ch
}

func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	fmt.Println(findWords(board, words)) // [eat oath]

	board2 := [][]byte{{'a', 'b'}, {'c', 'd'}}
	fmt.Println(findWords(board2, []string{"abcb"})) // []
}
