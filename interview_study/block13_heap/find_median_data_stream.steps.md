## Hint 1
Two heaps: max-heap `low` for the lower half, min-heap `high` for the upper half. Invariant: low.Len() == high.Len() or low.Len() == high.Len()+1.

```go
type MedianFinder struct {
    low  *MaxHeap // lower half
    high *MinHeap // upper half
}
```

## Hint 2
On AddNum, the cleanest pattern is: push into low unconditionally, then move its top to high. This guarantees max(low) <= min(high).

```go
heap.Push(m.low, n)
heap.Push(m.high, heap.Pop(m.low))
```

## Hint 3
That last step might over-shift, leaving high larger than low. Rebalance by moving back if needed.

```go
if m.high.Len() > m.low.Len() {
    heap.Push(m.low, heap.Pop(m.high))
}
```

## Hint 4
FindMedian is O(1). If sizes match, average both tops; otherwise low has one extra, so its top is the median. Cast to float64 before dividing.

```go
if m.low.Len() > m.high.Len() {
    return float64((*m.low)[0])
}
return (float64((*m.low)[0]) + float64((*m.high)[0])) / 2.0
```
