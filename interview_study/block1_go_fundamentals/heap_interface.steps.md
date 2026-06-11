## Hint 1
`container/heap` operates on any type that satisfies heap.Interface — you provide 5 methods, it gives you Push/Pop/Init.

```go
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
```

## Hint 2
Push and Pop operate on the BACKING SLICE (append/truncate at the end). The package's heap.Push/heap.Pop handle the sift logic.

```go
func (h *IntHeap) Push(x any)   { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
    old := *h
    n := len(old) - 1
    x := old[n]
    *h = old[:n]
    return x
}
```

## Hint 3
Always go through `heap.Push(&h, v)` / `heap.Pop(&h)` — never call `h.Push` directly, or you'll corrupt the heap invariant.

```go
h := &IntHeap{5, 3, 8, 1}
heap.Init(h)
heap.Push(h, 2)
fmt.Println(heap.Pop(h)) // 1 (min)
```
