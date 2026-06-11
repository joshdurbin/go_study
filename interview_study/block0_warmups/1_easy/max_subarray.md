PROBLEM: Maximum Subarray (Kadane's Algorithm)
==============================================
Given an integer array, find the contiguous subarray with the largest sum and
return that sum.

Why it's a warm-up: introduces "running best" DP intuition in its simplest form.
At each index, either extend the current run or restart from this element.

Example:
  [-2,1,-3,4,-1,2,1,-5,4] → 6  (the subarray [4,-1,2,1])
  [1]                     → 1
  [5,4,-1,7,8]            → 23
