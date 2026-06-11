## Hint 1
Water above bar i is `min(maxLeft, maxRight) - height[i]`. Naive: precompute leftMax[] and rightMax[] in two passes. O(n) time, O(n) space.

```go
n := len(h)
leftMax := make([]int, n)
rightMax := make([]int, n)
// fill both, then sum min(leftMax[i], rightMax[i]) - h[i]
```

## Hint 2
The two-pointer trick removes the O(n) space. Two pointers (left, right) with running lMax, rMax.

```go
left, right := 0, len(h)-1
lMax, rMax, total := 0, 0, 0
for left < right {
    // advance the smaller side
}
return total
```

## Hint 3
At each step, the side with the SMALLER current value is the binding wall — its running max determines water at that index.

```go
if h[left] < h[right] {
    // left is binding
    left++
} else {
    // right is binding
    right--
}
```

## Hint 4
Update or accrue: if current ≥ running max, update; otherwise, the difference is trapped water.

```go
if h[left] < h[right] {
    if h[left] >= lMax { lMax = h[left] } else { total += lMax - h[left] }
    left++
} else {
    if h[right] >= rMax { rMax = h[right] } else { total += rMax - h[right] }
    right--
}
```
