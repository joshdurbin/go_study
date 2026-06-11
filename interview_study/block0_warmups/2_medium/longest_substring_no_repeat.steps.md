## Hint 1
Variable-size sliding window. Right always advances; left jumps forward when a repeat appears inside the window.

```go
lastSeen := make(map[byte]int)
best, left := 0, 0
for right := 0; right < len(s); right++ {
    // if duplicate inside window, jump left forward
    // update best
}
return best
```

## Hint 2
Track each character's last-seen index. If we see a duplicate whose prior index is ≥ left, jump left to (idx + 1) — old duplicates outside the window are irrelevant.

```go
if idx, ok := lastSeen[s[right]]; ok && idx >= left {
    left = idx + 1
}
lastSeen[s[right]] = right
```

## Hint 3
Update best each step. Window length is `right - left + 1`.

```go
if right-left+1 > best {
    best = right - left + 1
}
```
