## Hint 1
You don't need to track which path you took — only the farthest index reachable from anywhere visited so far. If the loop ever steps past that frontier, you're stuck.

```go
farthest := 0
for i, x := range nums {
    // if i > farthest, return false
    // else extend farthest with i + x
}
```

## Hint 2
At each step, extend `farthest = max(farthest, i+x)`. Short-circuit the moment `farthest` covers the last index.

```go
if i > farthest {
    return false
}
if i+x > farthest {
    farthest = i + x
}
if farthest >= len(nums)-1 {
    return true
}
```

## Hint 3
Pitfall: don't iterate past indices that aren't actually reachable — the `i > farthest` guard must run BEFORE you use `nums[i]` to extend. Otherwise a `0` in an unreachable cell could mask the failure.

```go
for i, x := range nums {
    if i > farthest { return false }
    if i+x > farthest { farthest = i + x }
}
return true
```
