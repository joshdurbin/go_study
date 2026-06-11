## Hint 1
Cycle detection: if two pointers move at different speeds, they meet inside a cycle. If no cycle, fast reaches nil.

```go
slow, fast := head, head
for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
    // detect meeting
}
return false
```

## Hint 2
On every iteration, check if they collide.

```go
slow = slow.Next
fast = fast.Next.Next
if slow == fast { return true }
```

## Hint 3
To find the cycle's START (Floyd's), after they meet, reset one pointer to head; advance both by 1 until they meet again — that's the entry.

```go
p := head
for p != slow {
    p = p.Next
    slow = slow.Next
}
return p // cycle entry
```
