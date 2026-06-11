## Hint 1
O(1) extra space trick: three reversals. First normalize k.

```go
n := len(a)
if n == 0 { return }
k %= n
// reverse all, reverse first k, reverse rest
```

## Hint 2
Write a small in-place reverse helper.

```go
func reverse(a []int, i, j int) {
    for i < j {
        a[i], a[j] = a[j], a[i]
        i++; j--
    }
}
```

## Hint 3
Apply the three reversals — order matters.

```go
reverse(a, 0, n-1)
reverse(a, 0, k-1)
reverse(a, k, n-1)
```
