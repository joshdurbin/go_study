## Hint 1
Naive nested loop is O(n²). Single pass with a map of `value → index`: for each x, ask "have I already seen `target - x`?"

```go
seen := make(map[int]int, len(nums))
for i, x := range nums {
    // check map for target - x, otherwise record x → i
}
```

## Hint 2
If the complement exists, we're done. Otherwise record current and continue.

```go
for i, x := range nums {
    if j, ok := seen[target-x]; ok {
        return []int{j, i}
    }
    seen[x] = i
}
return nil
```
