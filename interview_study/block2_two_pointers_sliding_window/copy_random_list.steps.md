## Hint 1
Easy path first: two passes with a `map[*Node]*Node`. Pass 1 clones each node and records original→clone. Pass 2 walks again wiring `clone.Next` and `clone.Random` via the map.

```go
m := map[*Node]*Node{}
for c := head; c != nil; c = c.Next {
    m[c] = &Node{Val: c.Val}
}
for c := head; c != nil; c = c.Next {
    m[c].Next = m[c.Next]
    m[c].Random = m[c.Random]
}
return m[head]
```

## Hint 2
O(1) space trick — interleave clones into the original list so each clone sits right after its original. Now `cur.Next` IS the clone of `cur`, giving O(1) lookup with no map.

```go
for cur := head; cur != nil; cur = cur.Next.Next {
    clone := &Node{Val: cur.Val, Next: cur.Next}
    cur.Next = clone
}
```

## Hint 3
Wire Random pointers using the adjacency: `cur.Next` is `cur`'s clone, so `cur.Random.Next` is the clone of `cur.Random`. Then unzip by restoring each original's Next and chaining the clones.

```go
for cur := head; cur != nil; cur = cur.Next.Next {
    if cur.Random != nil {
        cur.Next.Random = cur.Random.Next
    }
}
// unzip: separate originals and clones back into two lists
```
