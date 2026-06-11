PROBLEM: Daily Temperatures
===========================
Given a slice `temps` of daily temperatures, return a slice `result` where
`result[i]` is the number of days until a warmer temperature. If no warmer day
exists, `result[i] = 0`.

Why it's medium: the brute-force O(n^2) double loop is obvious; recognizing
that a monotonic decreasing stack collapses it to O(n) is the leap. Each index
gets pushed once and popped once.

Example:
  temps  = [73, 74, 75, 71, 69, 72, 76, 73]
  result = [ 1,  1,  4,  2,  1,  1,  0,  0]

  temps  = [30, 40, 50, 60]
  result = [ 1,  1,  1,  0]
