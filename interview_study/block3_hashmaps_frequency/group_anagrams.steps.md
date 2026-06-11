## Hint 1
Anagrams share a canonical form. Sorting each word's chars is the simplest key.

```go
groups := make(map[string][]string)
for _, w := range words {
    key := canonical(w)
    groups[key] = append(groups[key], w)
}
```

## Hint 2
Sort-key is O(k log k) per word — convert to []byte, sort, back to string.

```go
func canonical(w string) string {
    b := []byte(w)
    sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
    return string(b)
}
```

## Hint 3
Faster: count-key with a [26]int frequency array → string. O(k) per word for ASCII. Mention both approaches in interviews.

```go
func canonical(w string) string {
    var cnt [26]int
    for _, c := range w { cnt[c-'a']++ }
    var b strings.Builder
    for _, n := range cnt { fmt.Fprintf(&b, "%d#", n) }
    return b.String()
}
```
