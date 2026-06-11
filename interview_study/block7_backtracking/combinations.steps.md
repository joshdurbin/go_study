## Hint 1
Same template as subsets but only record at depth k. Recurse with a `start` so combos stay sorted and unique.

```go
var backtrack func(start int)
backtrack = func(start int) {
    if len(path) == k { /* record copy */ return }
}
```

## Hint 2
The unpruned loop is `for i := start; i <= n; i++`. That works but explores dead ends — if you need 2 more elements and only 1 number remains, you'll never finish.

```go
for i := start; i <= n; i++ {
    path = append(path, i)
    backtrack(i + 1)
    path = path[:len(path)-1]
}
```

## Hint 3
Prune the upper bound: with `need = k - len(path)` slots left, the largest start that can still complete is `n - need + 1`. Cuts branches massively for large n.

```go
need := k - len(path)
for i := start; i <= n-need+1; i++ {
    path = append(path, i)
    backtrack(i + 1)
    path = path[:len(path)-1]
}
```
