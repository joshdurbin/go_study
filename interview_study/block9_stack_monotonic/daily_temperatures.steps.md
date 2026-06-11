## Hint 1
"Next greater" → monotonic stack. Keep a stack of **indices** (not values) so you can compute the day-distance on pop.

```go
result := make([]int, len(temps))
stack := make([]int, 0, len(temps)) // indices with decreasing temps
for i, t := range temps {
    // pop while current t beats the stack top
}
return result
```

## Hint 2
When today's temp exceeds the stack's top temp, that popped index has finally found its warmer day — record `i - top`.

```go
for len(stack) > 0 && temps[stack[len(stack)-1]] < t {
    top := stack[len(stack)-1]
    stack = stack[:len(stack)-1]
    result[top] = i - top
}
```

## Hint 3
After popping, push the current index. Indices left on the stack at the end have no warmer day — their `result` stays 0 from the zero-value initialization.

```go
stack = append(stack, i)
// indices still on the stack keep result[k] = 0 — no warmer day found
```
