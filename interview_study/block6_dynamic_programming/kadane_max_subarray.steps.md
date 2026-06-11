## Hint 1
Define the DP state precisely: `dp[i]` = max sum of a contiguous subarray ending exactly at index i (not "anywhere in nums[:i+1]"). The recurrence drops out: extend the previous run, or start fresh at nums[i].

```go
dp := make([]int, len(nums))
dp[0] = nums[0]
for i := 1; i < len(nums); i++ {
    // dp[i] = max(nums[i], dp[i-1] + nums[i])
}
// answer = max(dp...)
```

## Hint 2
Only `dp[i-1]` is read at each step, so collapse the array into a single rolling scalar `curr`. Track `best` separately as the running max — the answer is not `curr` at the end, it's the maximum value `curr` ever took.

```go
curr, best := nums[0], nums[0]
for _, x := range nums[1:] {
    if curr+x > x { curr += x } else { curr = x }
    if curr > best { best = curr }
}
return best
```

## Hint 3
The all-negative case is the trap. Initializing `best := 0` is wrong — for `[-3,-1,-4,-2]` it returns 0, but the largest contiguous sum is `-1` (single element). Seed both `curr` and `best` from `nums[0]` so a non-empty array always returns a real element. Handle empty input as a separate guard.

```go
if len(nums) == 0 {
    return 0 // document this edge case
}
curr, best := nums[0], nums[0] // never seed with 0
```
