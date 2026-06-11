PROBLEM: Word Search
====================
Given an m×n grid of characters and a word, return true if the word can be
constructed from letters of sequentially adjacent cells (horizontal/vertical
neighbors). The same cell may not be used more than once in a single path.

Why it's medium: grid DFS plus the in-place "mark and restore" trick. Naive
`visited[][]` works but allocating it per call is wasteful — temporarily
overwriting board[r][c] with a sentinel like '#' is the idiomatic move.

Example:
  board=[[A,B,C,E],[S,F,C,S],[A,D,E,E]]
    word="ABCCED" → true
    word="SEE"    → true
    word="ABCB"   → false
