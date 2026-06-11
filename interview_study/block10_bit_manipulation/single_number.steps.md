## Hint 1
Hashmap counting works but uses O(n) space. Two XOR identities solve it in O(1): `x ^ x = 0` and `x ^ 0 = x`. Fold XOR across the whole array — every pair cancels.

```go
result := 0
for _, x := range nums {
    // XOR into result
}
```

## Hint 2
XOR is commutative and associative so the duplicates don't need to be adjacent. Whatever pairs up vanishes; the loner survives.

```go
result := 0
for _, x := range nums {
    result ^= x
}
return result
```
