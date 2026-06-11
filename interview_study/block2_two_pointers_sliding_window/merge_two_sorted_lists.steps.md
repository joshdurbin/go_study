## Hint 1
Use a dummy head + tail pointer. Two-pointer walk on both lists, splice the smaller head each step.

```go
dummy := &Node{}
tail := dummy
for a != nil && b != nil {
    // splice the smaller head
    tail = tail.Next
}
return dummy.Next
```

## Hint 2
Compare values, attach the smaller, advance that side.

```go
if a.Val <= b.Val {
    tail.Next = a; a = a.Next
} else {
    tail.Next = b; b = b.Next
}
tail = tail.Next
```

## Hint 3
When one runs out, splice the rest of the other in O(1) — no need to walk it.

```go
if a != nil {
    tail.Next = a
} else {
    tail.Next = b
}
```
