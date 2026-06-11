PROBLEM: Trapping Rain Water
============================
Given an array of non-negative integers representing bar heights of width 1,
compute how much water can be trapped after raining.

Why it's hard: the leap from O(n²) (for each i, scan left/right for max) to
O(n) with two pointers is non-obvious. The water above bar i is
  min(maxLeft, maxRight) - height[i].
The two-pointer trick: whichever side has the smaller current max determines
the answer for that index — so advance that side.

Example:
  [0,1,0,2,1,0,1,3,2,1,2,1] → 6
  [4,2,0,3,2,5]             → 9
