PROBLEM: Counting Bits
======================
Given a non-negative integer n, return an array result of length n+1 where
result[i] is the number of 1-bits in the binary representation of i.
Solve in O(n) total — not O(n·k) by counting each one independently.

Why it's medium: the naive answer is "popcount each value" at O(n log n). The
slick answer reuses prior results: i>>1 is i with its lowest bit dropped, and
you've already computed it.

Example:
  n=2 → [0,1,1]            (00, 01, 10)
  n=5 → [0,1,1,2,1,2]      (000, 001, 010, 011, 100, 101)
