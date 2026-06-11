## Hint 1
"Next greater" → monotonic decreasing stack of indices. Initialize result with -1 so unmatched indices stay -1.

```go
n := len(nums)
result := make([]int, n)
for i := range result { result[i] = -1 }
stack := make([]int, 0, n)
```

## Hint 2
Circular trick: loop `i` from 0 to `2n-1` and look up `nums[i % n]`. The second pass lets early indices wrap around and find a greater element later.

```go
for i := 0; i < 2*n; i++ {
    v := nums[i%n]
    for len(stack) > 0 && nums[stack[len(stack)-1]] < v {
        // pop and resolve
    }
}
```

## Hint 3
Only push original indices (`i < n`) — the second pass exists only to **resolve** pending pops, not to add new candidates.

```go
top := stack[len(stack)-1]
stack = stack[:len(stack)-1]
result[top] = v
// ...
if i < n {
    stack = append(stack, i)
}
```
