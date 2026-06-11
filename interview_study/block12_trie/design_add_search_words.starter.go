//go:build ignore

package main

import "fmt"

// WordDictionary supports AddWord and Search where '.' in a search query
// matches any single letter.
// Build: trie under the hood, recursive DFS at wildcard positions.
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

// AddWord inserts word into the dictionary.
func (d *WordDictionary) AddWord(word string) {
	// TODO: implement
}

// Search returns true if any inserted word matches the pattern. '.' matches any letter.
func (d *WordDictionary) Search(word string) bool {
	// TODO: implement (recursive DFS — branch on '.')
	return false
}

func main() {
	d := NewWordDictionary()
	d.AddWord("bad")
	d.AddWord("dad")
	d.AddWord("mad")
	fmt.Println(d.Search("pad"))  // expect false
	fmt.Println(d.Search("bad"))  // expect true
	fmt.Println(d.Search(".ad"))  // expect true
	fmt.Println(d.Search("b.."))  // expect true
	fmt.Println(d.Search("...")) // expect true
	fmt.Println(d.Search("....")) // expect false
	fmt.Println(d.Search("b.d"))  // expect true
	fmt.Println(d.Search("b.z"))  // expect false
}
