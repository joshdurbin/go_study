PROBLEM: Single Number
======================
Given a non-empty array of ints where every element appears twice except for one,
find the single element. Solve in O(n) time and O(1) extra space.

Why it's easy: the XOR trick is the canonical bit-manipulation interview question.
x^x=0 and x^0=x, so XORing every element collapses all pairs and leaves the loner.

Example:
  [2,2,1]       → 1
  [4,1,2,1,2]   → 4
  [1]           → 1
