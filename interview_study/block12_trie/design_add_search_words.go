//go:build ignore

package main

import "fmt"

// WordDictionary: trie + DFS for '.' wildcards.
// AddWord: O(L). Search: O(L) when no wildcards; O(26^k * L) worst case
// where k is the number of '.' chars (rare in practice).

type node struct {
	children [26]*node
	isEnd    bool
}

type WordDictionary struct {
	root *node
}

func NewWordDictionary() *WordDictionary {
	return &WordDictionary{root: &node{}}
}

func (d *WordDictionary) AddWord(word string) {
	n := d.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if n.children[c] == nil {
			n.children[c] = &node{}
		}
		n = n.children[c]
	}
	n.isEnd = true
}

func (d *WordDictionary) Search(word string) bool {
	return dfs(d.root, word, 0)
}

func dfs(n *node, word string, i int) bool {
	if n == nil {
		return false
	}
	if i == len(word) {
		return n.isEnd
	}
	ch := word[i]
	if ch == '.' {
		for _, child := range n.children {
			if child != nil && dfs(child, word, i+1) {
				return true
			}
		}
		return false
	}
	return dfs(n.children[ch-'a'], word, i+1)
}

func main() {
	d := NewWordDictionary()
	d.AddWord("bad")
	d.AddWord("dad")
	d.AddWord("mad")
	fmt.Println(d.Search("pad"))  // false
	fmt.Println(d.Search("bad"))  // true
	fmt.Println(d.Search(".ad"))  // true
	fmt.Println(d.Search("b.."))  // true
	fmt.Println(d.Search("..."))  // true
	fmt.Println(d.Search("...."))  // false
	fmt.Println(d.Search("b.d"))  // true
	fmt.Println(d.Search("b.z"))  // false
}
