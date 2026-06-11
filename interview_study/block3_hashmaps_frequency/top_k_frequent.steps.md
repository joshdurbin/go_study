## Hint 1
First pass: frequency map. Then top-k by frequency — a sort gives O(n log n), a size-k min-heap gives O(n log k).

```go
freq := make(map[int]int)
for _, x := range nums { freq[x]++ }
// then top-k by value
```

## Hint 2
Min-heap of size k over (count, value). When size exceeds k, pop the smallest. What's left is the top-k.

```go
h := &countHeap{} // min-heap on count
for v, c := range freq {
    heap.Push(h, pair{count: c, val: v})
    if h.Len() > k { heap.Pop(h) }
}
```

## Hint 3
Bucket sort is faster: buckets[count] = []value, since count ≤ n. Walk from high to low and collect k.

```go
buckets := make([][]int, len(nums)+1)
for v, c := range freq { buckets[c] = append(buckets[c], v) }
out := []int{}
for i := len(buckets) - 1; i >= 0 && len(out) < k; i-- {
    out = append(out, buckets[i]...)
}
return out[:k]
```
