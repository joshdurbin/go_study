## Hint 1
Naive recursion recomputes the same subproblems exponentially. You only ever need the previous two values.

```go
func fib(n int) int {
    if n < 2 { return n }
    a, b := 0, 1
    // roll forward n-1 times
}
```

## Hint 2
Each iteration: `a, b = b, a+b`. After the loop, b is fib(n).

```go
for i := 2; i <= n; i++ {
    a, b = b, a+b
}
return b
```
