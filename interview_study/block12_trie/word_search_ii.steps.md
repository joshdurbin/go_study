## Hint 1
Build a trie of all dictionary words once. Storing the full word string on the terminal node is the slickest way to emit hits during DFS without rebuilding the path.

```go
type tnode struct {
    children [26]*tnode
    word     string // non-empty = a word terminates here
}
```

## Hint 2
One DFS launched from every board cell. Carry the current trie node as a parameter; the recursion only descends when the child exists — automatic pruning.

```go
for r := 0; r < m; r++ {
    for c := 0; c < n; c++ {
        dfs(board, r, c, root, &out)
    }
}
```

## Hint 3
Inside dfs: bail on visited ('#') or missing trie child. Step into the child node; if word != "", emit and blank it. Mark visited, recurse 4 directions, restore.

```go
ch := board[r][c]
if ch == '#' || n.children[ch-'a'] == nil { return }
n = n.children[ch-'a']
if n.word != "" {
    *out = append(*out, n.word)
    n.word = "" // dedupe
}
board[r][c] = '#'
// recurse 4 neighbors...
board[r][c] = ch
```

## Hint 4
Bounds-check inline before each recursive call — cheaper than a helper. The board is the visited set; no separate map needed.

```go
if r > 0          { dfs(board, r-1, c, n, out) }
if r+1 < len(board) { dfs(board, r+1, c, n, out) }
if c > 0          { dfs(board, r, c-1, n, out) }
if c+1 < len(board[0]) { dfs(board, r, c+1, n, out) }
```
