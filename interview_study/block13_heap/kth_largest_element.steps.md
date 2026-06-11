## Hint 1
Use a MIN-heap of size k. Counterintuitive: a min-heap is the right tool because the smallest of the top-k IS the k-th largest.

```go
h := &MinHeap{}
for _, n := range nums {
    heap.Push(h, n)
}
```

## Hint 2
Cap the heap at size k by popping after each push. This evicts the smallest "candidate" as larger values arrive.

```go
for _, n := range nums {
    heap.Push(h, n)
    if h.Len() > k {
        heap.Pop(h)
    }
}
```

## Hint 3
After processing all input, the heap contains exactly the top-k largest. Its root is the smallest of those — i.e. the k-th largest.

```go
return (*h)[0]
```
