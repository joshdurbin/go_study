PROBLEM: Merge K Sorted Linked Lists
====================================
Given an array of k sorted linked lists, merge them into one sorted list and
return its head.

Why it's medium: the naive merge-pairwise approach is O(nk). A min-heap keyed
by the current head of each list gives O(n log k) — pop the smallest, advance
that list, push the new head.

Example:
  Input:  [[1,4,5], [1,3,4], [2,6]]
  Output: [1,1,2,3,4,4,5,6]
