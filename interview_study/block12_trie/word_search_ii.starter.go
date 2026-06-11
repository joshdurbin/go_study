//go:build ignore

package main

import "fmt"

// tnode: trie node. Store the full word string on terminal nodes so DFS can
// emit hits without rebuilding the path.
type tnode struct {
	children [26]*tnode
	word     string
}

// buildTrie inserts every word; terminal node's word field holds the full word.
func buildTrie(words []string) *tnode {
	// TODO: implement
	return &tnode{}
}

// findWords returns every word in `words` reachable on `board` via
// 4-directional moves with no cell reused per word. Result has no duplicates.
//
// Strategy: build a trie of words once, DFS from each cell carrying a trie
// pointer. Mark visited cells with '#', restore on backtrack. Blank out
// node.word after first hit to dedupe.
func findWords(board [][]byte, words []string) []string {
	// TODO: implement
	return nil
}

func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	fmt.Println(findWords(board, words)) // expect [eat oath]

	board2 := [][]byte{{'a', 'b'}, {'c', 'd'}}
	fmt.Println(findWords(board2, []string{"abcb"})) // expect []
}
