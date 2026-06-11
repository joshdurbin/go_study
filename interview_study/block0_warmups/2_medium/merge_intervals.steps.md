## Hint 1
Any "overlap of ranges" problem starts with sorting by start. After that, overlap is decided by comparing only to the last emitted interval.

```go
sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0]
})
```

## Hint 2
Maintain a result list. For each interval, either extend the last one or start a new one.

```go
out := [][]int{intervals[0]}
for _, x := range intervals[1:] {
    last := out[len(out)-1]
    // if overlap → extend; else append
}
```

## Hint 3
Touching counts as overlap (`<=`, not `<`). Extend with `max(last.end, x.end)`.

```go
if x[0] <= last[1] {
    if x[1] > last[1] { last[1] = x[1] }
} else {
    out = append(out, x)
}
```
