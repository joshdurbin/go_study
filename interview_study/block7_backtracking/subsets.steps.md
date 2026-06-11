## Hint 1
Backtracking template: a recursive helper with a `start` index and a mutable `path`. Record the path at every node, not just leaves.

```go
res := [][]int{}
path := []int{}
var backtrack func(start int)
backtrack = func(start int) {
    // append a copy of path to res
}
```

## Hint 2
Critical: append a **copy** of path. Slices share backing arrays — appending the live slice means later mutations overwrite earlier subsets.

```go
cp := make([]int, len(path))
copy(cp, path)
res = append(res, cp)
```

## Hint 3
The `start` parameter prevents revisiting earlier elements (which would give permutations). Choose nums[i], recurse with i+1, then un-choose.

```go
for i := start; i < len(nums); i++ {
    path = append(path, nums[i])
    backtrack(i + 1)
    path = path[:len(path)-1]
}
```
