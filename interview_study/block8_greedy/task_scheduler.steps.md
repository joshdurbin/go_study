## Hint 1
The bottleneck is the most-frequent task. Count frequencies into a 26-slot array, then find `maxFreq` and how many distinct tasks tie at that frequency.

```go
var freq [26]int
for _, t := range tasks {
    freq[t-'A']++
}
// scan freq to get maxFreq and tiesAtMax
```

## Hint 2
Imagine the most-frequent task `M` runs `maxFreq` times. Between consecutive `M`s you need `n` slots. That's `(maxFreq-1)` frames of width `(n+1)`, then a final tail. Each task tied at `maxFreq` contributes 1 to that tail.

```go
frame := (maxFreq-1)*(n+1) + tiesAtMax
```

## Hint 3
Pitfall: when there are LOTS of distinct tasks, the frames fill with real work — no idles needed. In that case `len(tasks)` is the floor. Take the max.

```go
if len(tasks) > frame {
    return len(tasks)
}
return frame
```
