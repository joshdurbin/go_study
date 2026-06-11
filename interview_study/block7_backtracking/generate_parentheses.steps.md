## Hint 1
Don't generate-and-filter. Backtrack with two counters: `open` (how many '(' placed) and `closed` (how many ')' placed). Stop at length 2n.

```go
var backtrack func(open, closed int)
backtrack = func(open, closed int) {
    if len(cur) == 2*n { /* record */ return }
}
```

## Hint 2
Two constraints that keep every prefix valid: you can add '(' only while `open < n`; you can add ')' only while `closed < open` (otherwise the close has no opener).

```go
if open < n {
    cur = append(cur, '(')
    backtrack(open+1, closed)
    cur = cur[:len(cur)-1]
}
```

## Hint 3
Mirror the same pattern for ')'. Use a single byte slice and rewind it after each recurse — avoids allocating a fresh string per branch.

```go
if closed < open {
    cur = append(cur, ')')
    backtrack(open, closed+1)
    cur = cur[:len(cur)-1]
}
```
