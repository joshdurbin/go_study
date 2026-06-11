PROBLEM: Top K Frequent Elements
==================================
Given an integer array and an integer k, return the k most frequent elements.

Example:
  Input:  nums=[1,1,1,2,2,3], k=2  -> [1,2]
  Input:  nums=[1], k=1             -> [1]

Approach 1 (interview-safe): count with map, sort by freq. O(n log n)
Approach 2 (optimal): count with map, maintain a min-heap of size k. O(n log k)

The min-heap approach: push each element; if heap size exceeds k, pop the minimum.
After all elements, the heap contains the top k by frequency.
