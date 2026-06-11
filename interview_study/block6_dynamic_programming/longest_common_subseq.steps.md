## Hint 1
2-D string-alignment DP. State: dp[i][j] = LCS length between a[:i] and b[:j].

```go
m, n := len(a), len(b)
dp := make([][]int, m+1)
for i := range dp { dp[i] = make([]int, n+1) }
```

## Hint 2
Base cases: dp[0][*] = 0 and dp[*][0] = 0 (an empty prefix matches nothing). Already true since make zeros them — but say it.

```go
// dp[i][0] = 0 for all i
// dp[0][j] = 0 for all j
```

## Hint 3
On a character match, extend the diagonal. Otherwise take the max of dropping a char from either string.

```go
for i := 1; i <= m; i++ {
    for j := 1; j <= n; j++ {
        if a[i-1] == b[j-1] {
            dp[i][j] = dp[i-1][j-1] + 1
        } else {
            dp[i][j] = max(dp[i-1][j], dp[i][j-1])
        }
    }
}
return dp[m][n]
```

## Hint 4
This template generalizes: longest common substring (reset to 0 on mismatch, track running max), shortest common supersequence (m + n - LCS), edit distance (3 ops). Recognize the shape and the recurrence falls out.

```go
// Longest common SUBSTRING variant
if a[i-1] == b[j-1] {
    dp[i][j] = dp[i-1][j-1] + 1
    if dp[i][j] > best { best = dp[i][j] }
}
// else dp[i][j] stays 0
```
