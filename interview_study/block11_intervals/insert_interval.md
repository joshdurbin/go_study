PROBLEM: Insert Interval
========================
Given a sorted list of non-overlapping `[start, end]` intervals and a new
interval, insert it and merge any overlaps. Return the result still sorted.

Why it's medium: input is already sorted, so the trick is a clean three-phase
walk — append the strictly-before, fold the overlapping into newInterval, then
append the strictly-after. No re-sort, single pass.

Example:
  [[1,3],[6,9]], new=[2,5]                       → [[1,5],[6,9]]
  [[1,2],[3,5],[6,7],[8,10],[12,16]], new=[4,8]  → [[1,2],[3,10],[12,16]]
