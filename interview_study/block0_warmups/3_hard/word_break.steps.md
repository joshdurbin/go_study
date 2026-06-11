## Hint 1
Looks like backtracking but recurses on overlapping subproblems → DP. Let `dp[i]` = "can s[:i] be segmented?".

```go
set := make(map[string]struct{}, len(dict))
for _, w := range dict { set[w] = struct{}{} }
dp := make([]bool, len(s)+1)
dp[0] = true
```

## Hint 2
For each i, look back for any j where dp[j] is true AND s[j:i] is in the dictionary.

```go
for i := 1; i <= len(s); i++ {
    for j := 0; j < i; j++ {
        // check dp[j] and s[j:i] in set
    }
}
```

## Hint 3
Set dp[i] true and break the inner loop on first match — no need to keep scanning.

```go
if dp[j] {
    if _, ok := set[s[j:i]]; ok {
        dp[i] = true
        break
    }
}
```

## Hint 4
Wrap up: return dp[len(s)].

```go
return dp[len(s)]
```
