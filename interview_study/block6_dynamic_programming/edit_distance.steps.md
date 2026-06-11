## Hint 1
Three operations (insert, delete, substitute) → three recurrence branches. State: dp[i][j] over prefixes.

```go
m, n := len(a), len(b)
dp := make([][]int, m+1)
for i := range dp { dp[i] = make([]int, n+1) }
```

## Hint 2
Base cases are easy to forget: converting prefix-of-a to empty needs i deletions; empty to prefix-of-b needs j insertions.

```go
for i := 0; i <= m; i++ { dp[i][0] = i }
for j := 0; j <= n; j++ { dp[0][j] = j }
```

## Hint 3
On match: no operation needed, copy the diagonal.

```go
for i := 1; i <= m; i++ {
    for j := 1; j <= n; j++ {
        if a[i-1] == b[j-1] {
            dp[i][j] = dp[i-1][j-1]
        } else {
            // three operations ↓
        }
    }
}
```

## Hint 4
On mismatch: 1 + min(delete, insert, substitute). Name each branch — it prevents off-by-one errors.

```go
dp[i][j] = 1 + min3(
    dp[i-1][j],   // delete a[i-1]
    dp[i][j-1],   // insert b[j-1]
    dp[i-1][j-1], // substitute
)
return dp[m][n]
```
