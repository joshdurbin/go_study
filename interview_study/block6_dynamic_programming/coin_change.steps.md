## Hint 1
Unbounded knapsack-style DP. Let dp[i] = minimum coins to make amount i.

```go
dp := make([]int, amount+1)
// fill with sentinel "infinity" except dp[0]
```

## Hint 2
Use `amount+1` as the infinity sentinel (no real answer can exceed that). dp[0] = 0.

```go
const INF = math.MaxInt
for i := range dp { dp[i] = INF }
dp[0] = 0
```

## Hint 3
For each amount i, try every coin c ≤ i and take the minimum.

```go
for i := 1; i <= amount; i++ {
    for _, c := range coins {
        if c <= i && dp[i-c] != INF && dp[i-c]+1 < dp[i] {
            dp[i] = dp[i-c] + 1
        }
    }
}
```

## Hint 4
Return -1 when nothing works. Compare against the sentinel.

```go
if dp[amount] == INF { return -1 }
return dp[amount]
```
