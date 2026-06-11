## Hint 1
At each index, decide: extend the current run or restart from this element. Track best-ending-here (cur) and best-ever (best).

```go
cur, best := nums[0], nums[0]
for _, x := range nums[1:] {
    // cur = max(x, cur+x); best = max(best, cur)
}
return best
```

## Hint 2
"Restart" wins when adding x to the current run is worse than starting over at x. That's exactly `cur+x < x`, i.e., `cur < 0`.

```go
if cur+x < x { cur = x } else { cur += x }
if cur > best { best = cur }
```
