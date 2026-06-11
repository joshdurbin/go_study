## Hint 1
Without sorting, you'd have to compare every pair — O(n²). Sort by start time and the only conflict possible is with the immediately previous meeting.

```go
sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0]
})
```

## Hint 2
Scan adjacent pairs. A conflict means the current meeting starts strictly before the previous one ended. Touching (`==`) is fine — one room frees as the next begins.

```go
for i := 1; i < len(intervals); i++ {
    if intervals[i][0] < intervals[i-1][1] {
        return false
    }
}
return true
```
