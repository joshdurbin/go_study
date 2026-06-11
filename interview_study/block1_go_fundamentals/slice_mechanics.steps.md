## Hint 1
"In-place, preserving order" → two-pointer write head. One pointer scans, the other marks where to write.

```go
func removeDuplicates(a []int) []int {
    seen := make(map[int]bool, len(a))
    write := 0
    // scan a; advance write on each kept value
    return a[:write]
}
```

## Hint 2
For each value, if unseen, mark it and write to a[write], then advance write.

```go
for _, v := range a {
    if !seen[v] {
        seen[v] = true
        a[write] = v
        write++
    }
}
```

## Hint 3
The returned slice shares the same backing array — only the length shrinks. Recognize this "filter in place" pattern for any "remove X preserving order" problem.

```go
return a[:write]
```
