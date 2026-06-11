## Hint 1
Go strings are bytes, but multi-byte runes will get shredded by a byte swap. Convert to `[]rune` first.

```go
func reverse(s string) string {
    r := []rune(s)
    // two pointers swap inward
    return string(r)
}
```

## Hint 2
Opposite-ends swap, walk pointers toward each other.

```go
for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
    r[i], r[j] = r[j], r[i]
}
```
