## Hint 1
Track four walls. Each iteration consumes one edge and tightens the corresponding wall.

```go
top, bottom := 0, len(m)-1
left, right := 0, len(m[0])-1
out := []int{}
for top <= bottom && left <= right {
    // 4 edges, then tighten
}
```

## Hint 2
Top row L→R, then right column T→B. Tighten top++ and right-- as you go.

```go
for c := left; c <= right; c++ { out = append(out, m[top][c]) }
top++
for r := top; r <= bottom; r++ { out = append(out, m[r][right]) }
right--
```

## Hint 3
Bottom row R→L and left column B→T — but guard against single-row/column remainders or you'll re-visit.

```go
if top <= bottom {
    for c := right; c >= left; c-- { out = append(out, m[bottom][c]) }
    bottom--
}
if left <= right {
    for r := bottom; r >= top; r-- { out = append(out, m[r][left]) }
    left++
}
```
