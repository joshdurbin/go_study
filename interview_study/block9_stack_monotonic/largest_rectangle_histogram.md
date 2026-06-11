PROBLEM: Largest Rectangle in Histogram
=======================================
Given `heights` where each entry is the height of a bar of width 1, find the
area of the largest axis-aligned rectangle that fits within the histogram.

Why it's hard: O(n^2) brute force is straightforward; the O(n) monotonic-stack
solution requires three insights — (1) keep an **increasing** stack of indices,
(2) on each pop the width is `i - newStackTop - 1`, not `i - poppedIndex`, and
(3) a sentinel 0 at the end is needed to drain the stack.

Example:
  heights = [2, 1, 5, 6, 2, 3]
  answer  = 10   // bars at indices 2..3, height 5, width 2

  heights = [2, 4]
  answer  = 4    // single bar of height 4
