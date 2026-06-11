PROBLEM: Power of Two
=====================
Return true if n is a power of two (1, 2, 4, 8, 16, ...). Solve in O(1) without
loops or recursion.

Why it's easy: a power of two has exactly one bit set in binary. The identity
`n & (n-1) == 0` clears that single bit, so the result is zero iff n was a pure
power of two. Don't forget the `n > 0` guard.

Example:
  1  → true   (2^0)
  16 → true   (2^4)
  3  → false
  0  → false
  -4 → false
