## Hint 1
DFS from every cell. A recursive helper takes (r, c, k) where k is the next character index in word. Base case: k reached len(word) — we matched the whole word.

```go
var dfs func(r, c, k int) bool
dfs = func(r, c, k int) bool {
    if k == len(word) { return true }
    // bounds + mismatch check
}
```

## Hint 2
Bounds and mismatch as one early-exit. Then try all four neighbors with `k+1`. Short-circuit on the first success.

```go
if r < 0 || r >= m || c < 0 || c >= n || board[r][c] != word[k] {
    return false
}
// recurse on 4 neighbors
```

## Hint 3
The trick: don't allocate a visited[][]. Overwrite board[r][c] with '#' before recursing (so it can't match anything), then restore it after. This is the "mark and restore" idiom.

```go
saved := board[r][c]
board[r][c] = '#'
found := dfs(r+1, c, k+1) || dfs(r-1, c, k+1) ||
         dfs(r, c+1, k+1) || dfs(r, c-1, k+1)
board[r][c] = saved
return found
```
