## Hint 1
Two invariants. Feasibility: a solution exists iff `sum(gas) >= sum(cost)`. Location: any prefix that drives the tank negative cannot contain the start, so the answer must be after it.

```go
total, tank, start := 0, 0, 0
for i := range gas {
    diff := gas[i] - cost[i]
    // accumulate total and tank
}
```

## Hint 2
Walk once. If `tank` goes negative at index `i`, no station from `start..i` could have been the answer — reset `start = i+1`, `tank = 0`.

```go
total += diff
tank += diff
if tank < 0 {
    start = i + 1
    tank = 0
}
```

## Hint 3
At the end, the running `start` is correct only if `total >= 0`. Otherwise return -1. Don't be tempted to early-return inside the loop — the global feasibility check is the only gate.

```go
if total < 0 {
    return -1
}
return start
```
