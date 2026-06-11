PROBLEM: LRU Cache
==================
Implement a Least Recently Used cache with O(1) Get and Put. On capacity
overflow, evict the least recently used key.

Why it's hard: combines two data structures (hashmap + doubly linked list) to
achieve O(1) on both operations. Iconic system design / coding hybrid question.

API:
  c := NewLRU(2)
  c.Put(1, 1); c.Put(2, 2)
  c.Get(1)        → 1
  c.Put(3, 3)     // evicts key 2 (LRU)
  c.Get(2)        → -1
  c.Put(4, 4)     // evicts key 1
  c.Get(1)        → -1
  c.Get(3)        → 3
  c.Get(4)        → 4

Hint: head = most recent, tail = least recent. On Get/Put, move node to head.
On overflow, drop the node at tail.
