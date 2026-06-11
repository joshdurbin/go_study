## Hint 1
1-D DP state: "best answer for the first i elements" (or "ending at i"). Define dp and base cases.

```go
// Example: house robber
dp := make([]int, len(nums))
dp[0] = nums[0]
dp[1] = max(nums[0], nums[1])
```

## Hint 2
Recurrence: at i, either skip (carry dp[i-1]) or take (dp[i-2] + nums[i]).

```go
for i := 2; i < len(nums); i++ {
    dp[i] = max(dp[i-1], dp[i-2] + nums[i])
}
return dp[len(nums)-1]
```

## Hint 3
Validate base cases against the recurrence. Skip them and dp[2] reads dp[0] and dp[1] — both must be correct.

```go
// dp[0] = nums[0]              ← only one house
// dp[1] = max(nums[0], nums[1]) ← pick the bigger
```

## Hint 4
Space optimization: most 1-D DPs look back O(1) steps. Replace the array with two scalars rolled forward.

```go
prev2, prev1 := 0, 0
for _, x := range nums {
    prev2, prev1 = prev1, max(prev1, prev2 + x)
}
return prev1
```
