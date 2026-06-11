PROBLEM: Find Median from Data Stream
=====================================
Design a data structure that supports two operations on a stream of integers:
AddNum(n) inserts a value; FindMedian() returns the median of all values seen
so far. AddNum should be O(log n); FindMedian should be O(1).

Why it's hard: a single sorted structure makes one op cheap and the other
expensive. The insight is two heaps: a MAX-heap holding the lower half and a
MIN-heap holding the upper half. The median is the top of the larger heap, or
the average of the two tops when sizes match. The subtlety is the rebalance —
always push into low first, then sift the top to high, then re-sift if needed.

Example:
  AddNum(1), AddNum(2), FindMedian() -> 1.5
  AddNum(3), FindMedian()            -> 2
