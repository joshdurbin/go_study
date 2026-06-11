## Hint 1
Variable window: right always advances; left advances only while the window violates the constraint.

```go
left, best := 0, 0
for right := 0; right < len(a); right++ {
    // include a[right] in window state
    // while constraint violated: shrink from left
    // update best
}
return best
```

## Hint 2
Inner shrink uses a `for` loop (not an `if`) — multiple removals may be needed.

```go
for windowInvalid() {
    // remove a[left] from window state
    left++
}
```

## Hint 3
Each element enters and leaves at most once → O(n) amortized. The window's state (sum, frequency map, etc.) must be O(1) to update on entry and exit.

```go
if right - left + 1 > best {
    best = right - left + 1
}
```
