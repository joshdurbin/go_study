## Hint 1
Phase 1 is plain cycle detection: slow advances by 1, fast by 2. If fast hits nil there's no cycle; otherwise they meet inside it.

```go
slow, fast := head, head
for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
    if slow == fast { break }
}
```

## Hint 2
After they meet, reset one pointer to head. The algebra (distance head→entry equals distance meeting→entry, mod cycle length) means walking both at speed 1 lands them on the entry.

```go
p := head
for p != slow {
    p = p.Next
    slow = slow.Next
}
return p
```

## Hint 3
Edge case: fast reached nil before meeting → no cycle, return nil. Wrap phase 2 inside the meeting branch so the no-cycle path returns nil cleanly.

```go
if slow == fast {
    // phase 2 here, return entry
}
// fell out of loop → no cycle
return nil
```
