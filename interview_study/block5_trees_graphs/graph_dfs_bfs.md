PROBLEM: Number of Islands
===========================
Given an m x n grid of '1' (land) and '0' (water), count the islands.
An island is formed by adjacent (horizontal/vertical) land cells.

Example:
  Grid:                   Grid:
  1 1 1 1 0               1 1 0 0 0
  1 1 0 1 0               1 1 0 0 0
  1 1 0 0 0               0 0 1 0 0
  0 0 0 0 0               0 0 0 1 1
  -> 1 island             -> 3 islands

Pattern: DFS on an implicit grid graph
  - Each cell is a node; edges connect adjacent land cells
  - DFS from each unvisited '1', marking visited cells as '0' (sink them)
  - Each DFS initiation = one new island

Directions array trick: dirs := [][2]int{{1,0},{-1,0},{0,1},{0,-1}}

Time: O(m*n)  Space: O(m*n) stack depth worst case
