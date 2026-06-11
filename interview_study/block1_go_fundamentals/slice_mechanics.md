PROBLEM: Slice Mechanics Practice
==================================
Write a function that:
1. Takes a slice of ints
2. Removes all duplicates (preserve original order)
3. Returns the result without allocating a second slice (in-place, using two-pointer write head)

This exercises append mechanics, len vs cap awareness, and the write-head two-pointer pattern.

Example:
  Input:  [1, 2, 2, 3, 1, 4]
  Output: [1, 2, 3, 4]

Constraints:
  - O(n) time with a map for seen values
  - Modify the input slice in-place, return the trimmed slice
