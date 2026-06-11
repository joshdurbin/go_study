PROBLEM: Find the Duplicate Number (Floyd's Cycle Detection)
=============================================================
Given an array of n+1 integers where each integer is in [1, n],
exactly one number is duplicated. Find it.

Constraints:
  - O(1) extra space (no extra array/map)
  - Must not modify the input array
  - O(n) time

Key insight: treat the array as a linked list where index i points to nums[i].
The duplicate creates a cycle. Use Floyd's tortoise-and-hare to find the entry.

Example:
  Input:  [1,3,4,2,2]
  Output: 2

  Input:  [3,1,3,4,2]
  Output: 3
