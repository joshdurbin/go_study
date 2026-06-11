## Hint 1
Looping `n & 1` then `n >>= 1` works in 32 iterations. The Kernighan trick is better: `n & (n-1)` clears the lowest set bit in one step, so you only loop k times for k bits.

```go
count := 0
for n != 0 {
    // clear lowest set bit, increment count
}
```

## Hint 2
Subtracting 1 flips the lowest set bit to 0 and turns every lower 0 into 1; ANDing with the original wipes that whole tail.

```go
count := 0
for n != 0 {
    n &= n - 1
    count++
}
return count
```
