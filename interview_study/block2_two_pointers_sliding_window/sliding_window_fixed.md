PROBLEM: Maximum Sum Subarray of Size K
========================================
Given an array of integers and a number K, find the maximum sum
of any contiguous subarray of size exactly K.

Example:
  Input:  [2, 1, 5, 1, 3, 2], K=3
  Output: 9  (subarray [5,1,3])

  Input:  [2, 3, 4, 1, 5], K=2
  Output: 7  (subarray [3,4])

Pattern: fixed-size sliding window
  - Compute sum of first K elements
  - Slide right: add nums[i], subtract nums[i-K]
  - Track running max

Time: O(n)  Space: O(1)
