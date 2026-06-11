## Hint 1
Fixed window of k: maintain a running aggregate by adding the entering element and subtracting the leaving one — no recompute.

```go
sum := 0
for i := 0; i < k; i++ { sum += a[i] }
best := sum
// slide
```

## Hint 2
Slide one step at a time: `sum += a[i] - a[i-k]`.

```go
for i := k; i < len(a); i++ {
    sum += a[i] - a[i-k]
    if sum > best { best = sum }
}
return best
```

## Hint 3
Pattern generalizes beyond sums — any commutative aggregate (count of distinct via frequency map, etc.) can be updated O(1) per step.

```go
// frequency-map variant: on enter, freq[a[i]]++; on leave, freq[a[i-k]]--
// (and delete the key if it hits 0 to keep len(freq) meaningful)
```
