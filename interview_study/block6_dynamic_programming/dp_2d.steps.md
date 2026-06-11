## Hint 1
2-D DP: state needs two indices. Grid problems usually map cell (i,j) to dp[i][j].

```go
// Example: unique paths in a grid (right or down only)
dp := make([][]int, m)
for i := range dp { dp[i] = make([]int, n) }
```

## Hint 2
Initialize the first row and column directly — they have only one way in.

```go
for i := 0; i < m; i++ { dp[i][0] = 1 }
for j := 0; j < n; j++ { dp[0][j] = 1 }
```

## Hint 3
Recurrence pulls from dp[i-1][j] (came from above) and dp[i][j-1] (came from left).

```go
for i := 1; i < m; i++ {
    for j := 1; j < n; j++ {
        dp[i][j] = dp[i-1][j] + dp[i][j-1]
    }
}
return dp[m-1][n-1]
```

## Hint 4
Fill order matters: every cell you read must already be filled. Top-to-bottom, left-to-right is the standard for grid recurrences. Space can drop to one row if dp[i][j] depends only on dp[i-1][*] and dp[i][j-1].

```go
// 1-row optimization for unique paths
row := make([]int, n)
for j := range row { row[j] = 1 }
for i := 1; i < m; i++ {
    for j := 1; j < n; j++ {
        row[j] += row[j-1]
    }
}
return row[n-1]
```
