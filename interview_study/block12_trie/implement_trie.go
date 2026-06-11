//go:build ignore

package main

import "fmt"

// Trie: prefix tree over lowercase ASCII. Each node has a 26-slot child array
// and an isEnd flag marking a complete inserted word.
// Insert/Search/StartsWith are all O(L) where L is the word length.

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func NewTrie() *Trie { return &Trie{} }

func (t *Trie) Insert(word string) {
	node := t
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if node.children[c] == nil {
			node.children[c] = &Trie{}
		}
		node = node.children[c]
	}
	node.isEnd = true
}

// Search returns true only on exact-match insert.
func (t *Trie) Search(word string) bool {
	node := t.walk(word)
	return node != nil && node.isEnd
}

// StartsWith returns true if any inserted word has this prefix.
func (t *Trie) StartsWith(prefix string) bool {
	return t.walk(prefix) != nil
}

func (t *Trie) walk(s string) *Trie {
	node := t
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		if node.children[c] == nil {
			return nil
		}
		node = node.children[c]
	}
	return node
}

func main() {
	t := NewTrie()
	t.Insert("apple")
	fmt.Println(t.Search("apple"))   // true
	fmt.Println(t.Search("app"))     // false
	fmt.Println(t.StartsWith("app")) // true
	t.Insert("app")
	fmt.Println(t.Search("app"))     // true
	t.Insert("application")
	fmt.Println(t.StartsWith("appl")) // true
	fmt.Println(t.Search("apples"))   // false
	fmt.Println(t.StartsWith("bana")) // false
}
