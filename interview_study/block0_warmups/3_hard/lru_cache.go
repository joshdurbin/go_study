//go:build ignore

package main

import "fmt"

// LRU cache: hashmap (key → *node) + doubly-linked list (recency order).
// Sentinel head/tail nodes simplify edge cases (no nil checks on neighbors).

type node struct {
	key, val   int
	prev, next *node
}

type LRU struct {
	cap        int
	m          map[int]*node
	head, tail *node // sentinels: head.next = MRU, tail.prev = LRU
}

func NewLRU(cap int) *LRU {
	h, t := &node{}, &node{}
	h.next, t.prev = t, h
	return &LRU{cap: cap, m: make(map[int]*node, cap), head: h, tail: t}
}

func (c *LRU) Get(key int) int {
	n, ok := c.m[key]
	if !ok {
		return -1
	}
	c.moveToFront(n)
	return n.val
}

func (c *LRU) Put(key, val int) {
	if n, ok := c.m[key]; ok {
		n.val = val
		c.moveToFront(n)
		return
	}
	n := &node{key: key, val: val}
	c.m[key] = n
	c.addFront(n)
	if len(c.m) > c.cap {
		lru := c.tail.prev
		c.remove(lru)
		delete(c.m, lru.key)
	}
}

func (c *LRU) addFront(n *node) {
	n.prev, n.next = c.head, c.head.next
	c.head.next.prev = n
	c.head.next = n
}

func (c *LRU) remove(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (c *LRU) moveToFront(n *node) {
	c.remove(n)
	c.addFront(n)
}

func main() {
	c := NewLRU(2)
	c.Put(1, 1)
	c.Put(2, 2)
	fmt.Println(c.Get(1)) // 1
	c.Put(3, 3)           // evict 2
	fmt.Println(c.Get(2)) // -1
	c.Put(4, 4)           // evict 1
	fmt.Println(c.Get(1)) // -1
	fmt.Println(c.Get(3)) // 3
	fmt.Println(c.Get(4)) // 4
}
