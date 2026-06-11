## Hint 1
"0/1" = each item once. Pick: skip i, or take i (paying weight, gaining value).

```go
// 2-D form first to see the recurrence clearly
dp := make([][]int, n+1)
for i := range dp { dp[i] = make([]int, capacity+1) }
```

## Hint 2
2-D recurrence: dp[i][w] = max(skip, take if fits).

```go
for i := 1; i <= n; i++ {
    for w := 0; w <= capacity; w++ {
        dp[i][w] = dp[i-1][w] // skip
        if weights[i-1] <= w {
            v := dp[i-1][w - weights[i-1]] + values[i-1]
            if v > dp[i][w] { dp[i][w] = v }
        }
    }
}
```

## Hint 3
Compress to 1-D rolling array.

```go
dp := make([]int, capacity+1)
for i := 0; i < n; i++ {
    // inner loop ↓
}
```

## Hint 4
CRITICAL: iterate weight RIGHT-TO-LEFT in the inner loop. Left-to-right would re-use the same item (that's a different problem).

```go
for c := capacity; c >= weights[i]; c-- {
    if dp[c-weights[i]] + values[i] > dp[c] {
        dp[c] = dp[c-weights[i]] + values[i]
    }
}
return dp[capacity]
```
