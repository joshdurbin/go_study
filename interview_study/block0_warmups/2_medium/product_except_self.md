PROBLEM: Product of Array Except Self
=====================================
Given nums, return an array where out[i] = product of all elements except nums[i].
Solve in O(n) time WITHOUT using division.

Why it's medium: the "no division" twist forces a two-pass prefix/suffix product
trick. Hits prefix-array intuition without the usual sum framing.

Example:
  [1,2,3,4]   → [24,12,8,6]
  [-1,1,0,-3,3] → [0,0,9,0,0]
