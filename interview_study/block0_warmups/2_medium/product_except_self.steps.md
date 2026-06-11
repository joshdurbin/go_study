## Hint 1
No division allowed (zeros would break it anyway). Use prefix and suffix products: out[i] = (product left of i) × (product right of i).

```go
n := len(nums)
out := make([]int, n)
// pass 1: out[i] = product of everything LEFT of i
```

## Hint 2
Pass 1 left-to-right, accumulating the running prefix product:

```go
out[0] = 1
for i := 1; i < n; i++ {
    out[i] = out[i-1] * nums[i-1]
}
```

## Hint 3
Pass 2 right-to-left, multiplying in the running suffix:

```go
suffix := 1
for i := n - 1; i >= 0; i-- {
    out[i] *= suffix
    suffix *= nums[i]
}
return out
```
