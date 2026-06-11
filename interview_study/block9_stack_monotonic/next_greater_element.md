PROBLEM: Next Greater Element (Circular)
========================================
Given a **circular** integer slice `nums`, return a slice where each entry is
the next greater element when scanning forward (wrapping around). Use -1 if no
greater element exists in the entire circle.

Why it's medium: the monotonic-stack pattern is standard, but the circular
wraparound is the trick — iterate `2n` times using `i % n` so every index gets
a chance to see all candidates.

Example:
  nums   = [1, 2, 1]
  result = [2, -1, 2]   // last 1 wraps around to see 2

  nums   = [1, 2, 3, 4, 3]
  result = [2, 3, 4, -1, 4]
