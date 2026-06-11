## Hint 1
When answers form a monotonic predicate (false...false-true...true), binary search the ANSWER VALUE, not an index.

```go
lo, hi := minPossible, maxPossible
for lo < hi {
    mid := lo + (hi-lo)/2
    if feasible(mid) { hi = mid } else { lo = mid + 1 }
}
return lo
```

## Hint 2
feasible(x) must be monotone in x. Implement it as a normal O(n) function — usually a single pass.

```go
// Example: "can ship within D days with capacity x?"
func feasible(weights []int, D, x int) bool {
    days, load := 1, 0
    for _, w := range weights {
        if load + w > x { days++; load = 0 }
        load += w
    }
    return days <= D
}
```

## Hint 3
Total: O(n log(range)). Bounds matter — pick `lo = max(weights)` and `hi = sum(weights)` for shipping; tightening bounds avoids spurious iterations.

```go
lo, hi := 0, 0
for _, w := range weights {
    if w > lo { lo = w }
    hi += w
}
```
