## Hint 1
Step 1: find the end of the first half with slow/fast. Use `fast.Next != nil && fast.Next.Next != nil` so slow lands on the last node of the first half (for both even and odd lengths).

```go
slow, fast := head, head
for fast.Next != nil && fast.Next.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
}
```

## Hint 2
Step 2: detach and reverse the second half. Cut the first half by setting `slow.Next = nil`, then reverse `slow.Next` using the standard three-pointer pattern.

```go
var prev *ListNode
cur := slow.Next
slow.Next = nil
for cur != nil {
    next := cur.Next
    cur.Next = prev
    prev = cur
    cur = next
}
```

## Hint 3
Step 3: interleave. Walk `a` through the first half and `b` through the reversed second half, splicing one `b` node between each pair of `a` nodes. The second half is at most equal in length, so loop until `b == nil`.

```go
a, b := head, prev
for b != nil {
    aNext, bNext := a.Next, b.Next
    a.Next = b; b.Next = aNext
    a, b = aNext, bNext
}
```
