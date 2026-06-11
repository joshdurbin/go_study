//go:build ignore

package main

import "fmt"

// LRU is a Least-Recently-Used cache with O(1) Get and Put.
// On overflow, evict the least recently used key.
// Hint: hashmap (key → *node) + doubly-linked list (recency order) with
// sentinel head and tail nodes.
type LRU struct {
	// TODO: add fields
}

func NewLRU(cap int) *LRU {
	// TODO: implement
	return &LRU{}
}

// Get returns the value for key, or -1 if absent. Also marks key as most recent.
func (c *LRU) Get(key int) int {
	// TODO: implement
	return -1
}

// Put inserts or updates (key, val). On capacity overflow, evict the LRU entry.
func (c *LRU) Put(key, val int) {
	// TODO: implement
}

func main() {
	c := NewLRU(2)
	c.Put(1, 1)
	c.Put(2, 2)
	fmt.Println(c.Get(1)) // expect 1
	c.Put(3, 3)           // evicts 2
	fmt.Println(c.Get(2)) // expect -1
	c.Put(4, 4)           // evicts 1
	fmt.Println(c.Get(1)) // expect -1
	fmt.Println(c.Get(3)) // expect 3
	fmt.Println(c.Get(4)) // expect 4
}
