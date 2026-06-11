PROBLEM: Merge Intervals
========================
Given a list of [start, end] intervals, merge any that overlap and return the
result in sorted order.

Why it's medium: interview staple. Sort by start, then sweep — if the next
interval starts before the current end, extend; otherwise emit and move on.

Example:
  [[1,3],[2,6],[8,10],[15,18]] → [[1,6],[8,10],[15,18]]
  [[1,4],[4,5]]                → [[1,5]]   (touching counts as overlap)
