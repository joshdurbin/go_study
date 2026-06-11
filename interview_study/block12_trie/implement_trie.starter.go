//go:build ignore

package main

import "fmt"

// Trie is a prefix tree over lowercase ASCII.
// Each node has 26 child slots (one per letter) and an isEnd flag marking
// the last character of an inserted word. All ops are O(L) in word length.
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func NewTrie() *Trie { return &Trie{} }

// Insert adds word to the trie.
func (t *Trie) Insert(word string) {
	// TODO: implement
}

// Search returns true iff word was previously inserted (exact match).
func (t *Trie) Search(word string) bool {
	// TODO: implement
	return false
}

// StartsWith returns true iff any inserted word has prefix.
func (t *Trie) StartsWith(prefix string) bool {
	// TODO: implement
	return false
}

func main() {
	t := NewTrie()
	t.Insert("apple")
	fmt.Println(t.Search("apple"))   // expect true
	fmt.Println(t.Search("app"))     // expect false
	fmt.Println(t.StartsWith("app")) // expect true
	t.Insert("app")
	fmt.Println(t.Search("app")) // expect true
	t.Insert("application")
	fmt.Println(t.StartsWith("appl")) // expect true
	fmt.Println(t.Search("apples"))   // expect false
	fmt.Println(t.StartsWith("bana")) // expect false
}
