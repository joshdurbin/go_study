PROBLEM: Jump Game
==================
Given `nums` where `nums[i]` is the maximum jump length from index `i`, return
true if you can reach the last index starting from index 0.

Why it's medium: the canonical "greedy beats DP" problem. O(n) DP works but is
wasteful — tracking the single value "farthest reachable" collapses it to O(1) space.

Example:
  nums=[2,3,1,1,4] → true   (0 → 1 → 4)
  nums=[3,2,1,0,4] → false  (stuck at index 3)
  nums=[0]         → true   (already there)
