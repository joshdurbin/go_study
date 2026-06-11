## Hint 1
Seed the heap with the head of every non-nil list. The heap holds at most k nodes at any time.

```go
h := &nodeHeap{}
for _, head := range lists {
    if head != nil {
        heap.Push(h, head)
    }
}
```

## Hint 2
Use a dummy head + tail pointer to build the result. Pop the smallest, append, and push that node's Next if it exists.

```go
dummy := &ListNode{}
tail := dummy
for h.Len() > 0 {
    node := heap.Pop(h).(*ListNode)
    tail.Next = node
    tail = node
}
```

## Hint 3
The trick that keeps it O(n log k): only the current head of each list is in the heap at one time. After popping, push that node's successor.

```go
if node.Next != nil {
    heap.Push(h, node.Next)
}
return dummy.Next
```
