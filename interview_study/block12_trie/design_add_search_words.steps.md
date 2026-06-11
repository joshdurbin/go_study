## Hint 1
Standard trie node + AddWord. The wildcard logic lives entirely inside Search — insertion is unchanged.

```go
type node struct {
    children [26]*node
    isEnd    bool
}
```

## Hint 2
Search needs recursion: track the current node and the position in word. Base case: index equals len(word) → return isEnd.

```go
func dfs(n *node, word string, i int) bool {
    if n == nil { return false }
    if i == len(word) { return n.isEnd }
    // ... handle word[i]
}
```

## Hint 3
On a regular char, descend the one matching child. On '.', try every non-nil child and short-circuit on the first hit.

```go
ch := word[i]
if ch == '.' {
    for _, c := range n.children {
        if c != nil && dfs(c, word, i+1) { return true }
    }
    return false
}
return dfs(n.children[ch-'a'], word, i+1)
```
