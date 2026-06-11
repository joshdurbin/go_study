## Hint 1
Powers of two have exactly one bit set: `1, 10, 100, 1000, ...`. Subtracting 1 flips that bit off and lights every lower bit. AND-ing the two gives zero — that's the test.

```go
// what does n & (n-1) equal when n is a power of two?
// what about when n is 0 or negative?
```

## Hint 2
Combine the bit test with a positivity guard — 0 and negatives shouldn't return true even though they may satisfy the AND.

```go
return n > 0 && n&(n-1) == 0
```
