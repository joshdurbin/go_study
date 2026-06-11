## Hint 1
The greedy choice is "earliest end time". Sort by `end`, not by `start`. Picking the shortest-ending interval leaves the most runway for whatever comes next.

```go
sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][1] < intervals[j][1]
})
```

## Hint 2
Sweep once. Track `end` of the last kept interval. If the next one starts at or after `end`, keep it; otherwise it must be removed.

```go
end := intervals[0][1]
removals := 0
for _, x := range intervals[1:] {
    // compare x[0] to end
}
```

## Hint 3
Touching intervals (`x.start == end`) do not overlap — use `>=`. Don't update `end` on a removal: you're keeping the earlier (smaller-end) interval.

```go
if x[0] >= end {
    end = x[1]
} else {
    removals++
}
return removals
```
