PROBLEM: Median of Two Sorted Arrays
====================================
Given two sorted arrays, return the median of the merged set in O(log(min(m,n))).

Why it's hard: requires binary search over a PARTITION position, not values.
You're looking for the split where everything on the left ≤ everything on the
right across both arrays. Classic interview "hard" — be ready to bail to the
O(m+n) merge-and-pick approach if the binary search is too risky to nail live.

Example:
  [1,3], [2]   → 2.0
  [1,2], [3,4] → 2.5
