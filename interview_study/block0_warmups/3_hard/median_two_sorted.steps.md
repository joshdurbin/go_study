## Hint 1
O(m+n) merge works but doesn't hit the target. To beat it, binary search a PARTITION position — not a value.

```go
if len(a) > len(b) { a, b = b, a } // search the shorter
m, n := len(a), len(b)
half := (m + n + 1) / 2
lo, hi := 0, m
```

## Hint 2
For partition i in a (and j = half - i in b), define the four boundary values, using sentinels at the array edges.

```go
aLeft  := math.MinInt; if i > 0 { aLeft  = a[i-1] }
aRight := math.MaxInt; if i < m { aRight = a[i]   }
bLeft  := math.MinInt; if j > 0 { bLeft  = b[j-1] }
bRight := math.MaxInt; if j < n { bRight = b[j]   }
```

## Hint 3
Partition is valid when both left maxes ≤ both right mins. Otherwise tighten lo/hi.

```go
for lo <= hi {
    i := (lo + hi) / 2
    j := half - i
    // compute aLeft, aRight, bLeft, bRight
    if aLeft <= bRight && bLeft <= aRight {
        // valid — compute median
    } else if aLeft > bRight {
        hi = i - 1
    } else {
        lo = i + 1
    }
}
```

## Hint 4
Odd total → max of left sides. Even total → average of max-left and min-right.

```go
if (m+n)%2 == 1 {
    return float64(max(aLeft, bLeft))
}
return float64(max(aLeft, bLeft)+min(aRight, bRight)) / 2.0
```
