## Hint 1
Three pointers: prev (already-reversed tail), cur (current), next (saved before rewiring).

```go
var prev *ListNode
cur := head
for cur != nil {
    // save next, point cur.Next at prev, advance
}
return prev
```

## Hint 2
The order inside the loop matters: save first, rewire, then advance.

```go
next := cur.Next
cur.Next = prev
prev = cur
cur = next
```

## Hint 3
When cur reaches nil, prev is the new head. Recursive version exists but uses O(n) stack — iterative is the interview default.

```go
return prev
```
