PROBLEM: Kth Largest Element in an Array
========================================
Given an unsorted integer array and an integer k, return the k-th largest
element. (k=1 means the maximum.)

Why it's medium: sorting is O(n log n) — easy but wasteful. A min-heap of size
k gives O(n log k) and only O(k) memory. Counterintuitive at first: you use a
MIN-heap to find the LARGEST elements, because the heap holds the top k and
the smallest of those is exactly the k-th largest.

Example:
  Input:  nums = [3,2,1,5,6,4], k = 2
  Output: 5
