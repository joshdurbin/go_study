## Hint 1
Each node owns a 26-slot child array plus an isEnd flag. The root is an empty Trie value; nothing special.

```go
type Trie struct {
    children [26]*Trie
    isEnd    bool
}
```

## Hint 2
Insert walks one node per character, creating empty children as it goes, then sets isEnd on the final node.

```go
func (t *Trie) Insert(word string) {
    node := t
    for i := 0; i < len(word); i++ {
        c := word[i] - 'a'
        if node.children[c] == nil {
            node.children[c] = &Trie{}
        }
        node = node.children[c]
    }
    node.isEnd = true
}
```

## Hint 3
Search and StartsWith share a walk helper. The only difference: Search also requires isEnd; StartsWith just needs the node to exist.

```go
func (t *Trie) walk(s string) *Trie {
    node := t
    for i := 0; i < len(s); i++ {
        c := s[i] - 'a'
        if node.children[c] == nil { return nil }
        node = node.children[c]
    }
    return node
}
func (t *Trie) Search(w string) bool     { n := t.walk(w); return n != nil && n.isEnd }
func (t *Trie) StartsWith(p string) bool { return t.walk(p) != nil }
```
