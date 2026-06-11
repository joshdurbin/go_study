PROBLEM: Spiral Matrix Traversal
================================
Given an m x n matrix, return all elements in spiral (clockwise) order, starting
from the top-left.

Why it's medium: boundary-shrinking pattern. Track top/bottom/left/right walls
and tighten each after traversing its edge.

Example:
  [[1,2,3],
   [4,5,6],
   [7,8,9]]  →  [1,2,3,6,9,8,7,4,5]

  [[1,2,3,4],
   [5,6,7,8],
   [9,10,11,12]] → [1,2,3,4,8,12,11,10,9,5,6,7]
