PROBLEM: Hamming Weight
=======================
Given a uint32, return the number of 1-bits in its binary representation
(also called the population count or Hamming weight).

Why it's easy: the obvious solution checks all 32 bits one by one. The slick
solution uses `n & (n-1)` to clear the lowest set bit each step, so the loop
runs only as many times as there are 1-bits.

Example:
  0b...01011 → 3   (bits set: positions 0, 1, 3)
  0b...10000000 → 1
  0b...11111101 → 31
