## Hint 1
Sorted array + "find pair summing to target" → opposite-end two pointers.

```go
left, right := 0, len(a)-1
for left < right {
    sum := a[left] + a[right]
    // narrow based on sum vs target
}
return nil
```

## Hint 2
If sum equals target → found. Less → need larger → left++. More → need smaller → right--.

```go
switch {
case sum == target: return []int{left, right}
case sum < target:  left++
default:            right--
}
```

## Hint 3
Generalizes to 3-sum: fix one element, run opposite-end pointers on the rest. Sorting first is the unlock.

```go
sort.Ints(nums)
for i := 0; i < len(nums)-2; i++ {
    // l, r = i+1, len(nums)-1; same opposite-end loop targeting -nums[i]
}
```
