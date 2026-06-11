## Hint 1
Process meetings in start-time order. The heap holds the end times of currently-in-use rooms.

```go
sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0]
})
h := &MinHeap{}
```

## Hint 2
For each new meeting, peek the room ending soonest. If it's already done by the new start, reuse it (pop the old end).

```go
for _, m := range intervals {
    if h.Len() > 0 && (*h)[0] <= m[0] {
        heap.Pop(h)
    }
    // then push current end
}
```

## Hint 3
Always push the current meeting's end — whether reusing a room or allocating a new one. The peak heap size = peak concurrency = answer.

```go
heap.Push(h, m[1])
// after loop:
return h.Len()
```
