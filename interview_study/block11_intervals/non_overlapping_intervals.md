PROBLEM: Non-overlapping Intervals
==================================
Given a list of `[start, end]` intervals, return the minimum number of intervals
to remove so the rest are non-overlapping.

Why it's medium: classic interval scheduling. Greedy by earliest end time is
optimal — the interval that frees the room soonest leaves the most room for
future picks. Sort by end, then count whatever you have to drop.

Example:
  [[1,2],[2,3],[3,4],[1,3]] → 1   (drop [1,3])
  [[1,2],[1,2],[1,2]]       → 2   (keep one of three)
  [[1,2],[2,3]]             → 0   (touching is OK)
