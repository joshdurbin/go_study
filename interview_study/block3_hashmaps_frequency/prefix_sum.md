PROBLEM: Subarray Sum Equals K
================================
Given an integer array and an integer k, return the total number of
contiguous subarrays whose sum equals k.

Example:
  Input:  nums=[1,1,1], k=2    -> 2
  Input:  nums=[1,2,3], k=3    -> 2  ([1,2] and [3])
  Input:  nums=[-1,-1,1], k=0  -> 1

Key insight:
  If prefix[j] - prefix[i] == k, then subarray (i..j] sums to k.
  Rearranged: we need to have seen prefix[j] - k as a prior prefix sum.
  Build prefix sum as we go, store counts in a map.
  Initialize map with {0: 1} to handle subarrays starting at index 0.

Time: O(n)  Space: O(n)
