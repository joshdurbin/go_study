PROBLEM: Permutations
=====================
Given an array of distinct integers `nums`, return all possible permutations.
There are n! of them.

Why it's medium: two valid approaches — swap-in-place (O(1) extra space per
frame) vs. a `used[]` boolean set (more code, easier to extend to duplicates).
Knowing both signals depth.

Example:
  nums=[1,2,3] → [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,2,1],[3,1,2]]
  nums=[0,1]   → [[0,1],[1,0]]
