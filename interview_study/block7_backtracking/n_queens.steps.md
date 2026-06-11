## Hint 1
Place queens row by row — one queen per row by construction eliminates row conflicts. State you need: which column each row's queen sits in, plus three "attack" sets.

```go
queens := make([]int, n)   // queens[r] = column
cols := make([]bool, n)
diag := make([]bool, 2*n-1) // \ diagonal: r - c constant
anti := make([]bool, 2*n-1) // / diagonal: r + c constant
```

## Hint 2
Diagonals have nice invariants: on a `\` diagonal `r - c` is constant; on a `/` diagonal `r + c` is constant. Shift `r - c` by `n-1` so it indexes a 0-based array.

```go
d := r - c + (n - 1)
a := r + c
if cols[c] || diag[d] || anti[a] {
    continue
}
```

## Hint 3
Main loop: for each column in the current row, if the cell isn't under attack, mark all three sets, recurse to the next row, then un-mark on return. Classic backtracking.

```go
queens[r] = c
cols[c], diag[d], anti[a] = true, true, true
backtrack(r + 1)
cols[c], diag[d], anti[a] = false, false, false
```

## Hint 4
At r == n you've placed all queens — translate queens[] into the []string board format. Use strings.Repeat to build each row from dots with a single 'Q' inserted.

```go
if r == n {
    board := make([]string, n)
    for i := 0; i < n; i++ {
        board[i] = strings.Repeat(".", queens[i]) + "Q" + strings.Repeat(".", n-queens[i]-1)
    }
    res = append(res, board)
    return
}
```
