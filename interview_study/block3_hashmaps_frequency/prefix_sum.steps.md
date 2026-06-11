## Hint 1
Subarray sum k → prefix sums + hashmap. Sum of a[i..j] = P[j+1] - P[i]. So "sum = k" means a prior prefix = current - k.

```go
counts := map[int]int{0: 1} // empty prefix
running, ans := 0, 0
for _, x := range a {
    running += x
    // count prior prefixes equal to running - k
    counts[running]++
}
return ans
```

## Hint 2
At each step, ans += count[running - k] BEFORE recording the current prefix.

```go
if c, ok := counts[running-k]; ok {
    ans += c
}
counts[running]++
```

## Hint 3
The seed `counts[0] = 1` covers subarrays that start at index 0 (their prior prefix is the empty one).

```go
counts := map[int]int{0: 1}
```
