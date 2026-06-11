## Hint 1
Swap-in-place is the cleanest. Recurse over a `first` position; each call decides what value lands at position `first` by swapping it in from somewhere in nums[first:].

```go
var backtrack func(first int)
backtrack = func(first int) {
    if first == len(nums) {
        // record a copy of nums
    }
}
```

## Hint 2
When `first == n`, you have a complete permutation — copy it (slices share memory) into the result.

```go
if first == n {
    cp := make([]int, n)
    copy(cp, nums)
    res = append(res, cp)
    return
}
```

## Hint 3
For each i in [first, n), swap nums[first] with nums[i], recurse on first+1, then swap back to restore state for the next iteration. The swap-back is what makes it backtracking.

```go
for i := first; i < n; i++ {
    nums[first], nums[i] = nums[i], nums[first]
    backtrack(first + 1)
    nums[first], nums[i] = nums[i], nums[first]
}
```
