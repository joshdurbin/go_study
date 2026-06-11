## Hint 1
Input is already sorted, so resist re-sorting. Walk in three phases: before, overlapping, after. "Before" means `intervals[i].end < newInterval.start`.

```go
out := make([][]int, 0, len(intervals)+1)
i, n := 0, len(intervals)
for i < n && intervals[i][1] < newInterval[0] {
    out = append(out, intervals[i])
    i++
}
```

## Hint 2
An interval overlaps newInterval while `intervals[i].start <= newInterval.end`. Fold each into newInterval by widening its bounds.

```go
for i < n && intervals[i][0] <= newInterval[1] {
    if intervals[i][0] < newInterval[0] { newInterval[0] = intervals[i][0] }
    if intervals[i][1] > newInterval[1] { newInterval[1] = intervals[i][1] }
    i++
}
out = append(out, newInterval)
```

## Hint 3
Emit the merged newInterval once, then drain the rest. Common mistakes: appending newInterval inside the overlap loop, or using `<` instead of `<=` for the touching case.

```go
for i < n {
    out = append(out, intervals[i])
    i++
}
return out
```
