## Hint 1
The brute approach popcounts each value independently — O(n log n). To hit O(n), notice that `i >> 1` is just `i` with its lowest bit shifted away. You already computed it.

```go
result := make([]int, n+1)
// result[0] is 0 by zero-value default
for i := 1; i <= n; i++ {
    // recurrence using result[i>>1]
}
```

## Hint 2
Build the recurrence: `popcount(i) = popcount(i>>1) + (i & 1)`. The lowest bit either contributes 1 or doesn't; the rest matches a smaller index you've already solved.

```go
for i := 1; i <= n; i++ {
    result[i] = result[i>>1] + (i & 1)
}
```

## Hint 3
Because `i>>1 < i`, each `result[i]` reads only entries already filled — single forward pass, no recursion. Return the slice.

```go
result := make([]int, n+1)
for i := 1; i <= n; i++ {
    result[i] = result[i>>1] + (i & 1)
}
return result
```
