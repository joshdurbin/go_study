PROBLEM: Meeting Rooms II
=========================
Given an array of meeting time intervals [start, end), return the minimum
number of conference rooms required so that no two meetings collide.

Why it's medium: the trick is recognizing that "minimum rooms" = "peak
concurrent meetings". Sort by start, then keep a min-heap of in-use room end
times. When a new meeting starts, if the room ending soonest is already free,
reuse it (pop); always push the new end. Final heap size = answer.

Example:
  Input:  [[0,30],[5,10],[15,20]]
  Output: 2
