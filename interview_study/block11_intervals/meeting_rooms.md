PROBLEM: Meeting Rooms
======================
Given an array of meeting `[start, end]` intervals, determine if a person could
attend all of them — i.e. no two meetings overlap.

Why it's easy: sort by start time, then check each adjacent pair. If any meeting
starts before the previous one ended, it's impossible. Touching is allowed
(end == next.start).

Example:
  [[0,30],[5,10],[15,20]] → false  ([5,10] starts during [0,30])
  [[7,10],[2,4]]          → true
  [[1,5],[5,8]]           → true   (touching OK)
