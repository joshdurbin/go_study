PROBLEM: N-Queens
=================
Place n queens on an n×n chessboard so that no two attack each other (no
shared row, column, or diagonal). Return all distinct solutions. Each
solution is an []string where 'Q' marks a queen and '.' an empty square.

Why it's hard: pure backtracking is O(n^n); the win is the conflict
representation. Iterate by row (one queen per row by construction), then
track three boolean arrays for columns, `\` diagonals (r-c is constant
along these), and `/` diagonals (r+c is constant). Constant-time conflict
check turns the search from intractable into instant for n up to ~12.

Example:
  n=4 → 2 solutions:
    .Q..      ..Q.
    ...Q      Q...
    Q...      ...Q
    ..Q.      .Q..
  n=1 → [["Q"]]
