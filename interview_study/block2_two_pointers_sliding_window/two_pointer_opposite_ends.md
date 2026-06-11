PROBLEM: Container With Most Water
===================================
Given n non-negative integers representing heights of vertical lines at positions 0..n-1,
find two lines that together with the x-axis form a container holding the most water.

Example:
  Input:  [1,8,6,2,5,4,8,3,7]
  Output: 49  (lines at index 1 and 8, width=7, min height=7)

Pattern: opposite-end two pointer.
  - Start with widest possible container (lo=0, hi=n-1)
  - Move the shorter side inward (keeping it can only make things worse)
  - Track max area at each step

Time: O(n)  Space: O(1)
