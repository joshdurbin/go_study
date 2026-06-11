## Hint 1
O(1) Get + Put + eviction needs a hashmap (lookup) + doubly-linked list (move-to-front, evict-tail).

```go
type node struct {
    key, val   int
    prev, next *node
}
type LRU struct {
    cap        int
    m          map[int]*node
    head, tail *node // sentinels
}
```

## Hint 2
Sentinel head and tail nodes mean every real node always has neighbors — no nil checks in splice operations.

```go
func NewLRU(cap int) *LRU {
    h, t := &node{}, &node{}
    h.next, t.prev = t, h
    return &LRU{cap: cap, m: make(map[int]*node, cap), head: h, tail: t}
}
```

## Hint 3
Three list helpers cover every operation: addFront, remove, moveToFront.

```go
func (c *LRU) addFront(n *node) {
    n.prev, n.next = c.head, c.head.next
    c.head.next.prev = n
    c.head.next = n
}
func (c *LRU) remove(n *node) {
    n.prev.next = n.next
    n.next.prev = n.prev
}
func (c *LRU) moveToFront(n *node) { c.remove(n); c.addFront(n) }
```

## Hint 4
Get: lookup, move to front, return value. Put: update or insert; if over capacity, drop the node before tail.

```go
func (c *LRU) Get(k int) int {
    n, ok := c.m[k]
    if !ok { return -1 }
    c.moveToFront(n); return n.val
}
func (c *LRU) Put(k, v int) {
    if n, ok := c.m[k]; ok { n.val = v; c.moveToFront(n); return }
    n := &node{key: k, val: v}
    c.m[k] = n; c.addFront(n)
    if len(c.m) > c.cap {
        lru := c.tail.prev
        c.remove(lru); delete(c.m, lru.key)
    }
}
```
