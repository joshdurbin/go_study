## Hint 1
Pick one template and stick to it. Closed interval [lo, hi] is the most forgiving.

```go
func search(a []int, target int) int {
    lo, hi := 0, len(a)-1
    for lo <= hi {
        mid := lo + (hi-lo)/2
        // narrow
    }
    return -1
}
```

## Hint 2
Use `lo + (hi-lo)/2` not `(lo+hi)/2` — the latter overflows on huge ranges.

```go
mid := lo + (hi-lo)/2
switch {
case a[mid] == target: return mid
case a[mid] < target:  lo = mid + 1
default:               hi = mid - 1
}
```

## Hint 3
For "first index where P is true" (boundary search), narrow toward the boundary and return lo at the end.

```go
lo, hi := 0, n
for lo < hi {
    mid := lo + (hi-lo)/2
    if pred(mid) { hi = mid } else { lo = mid + 1 }
}
return lo
```
