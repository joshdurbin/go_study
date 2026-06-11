PROBLEM: Gas Station
====================
`gas[i]` is fuel at station `i`; `cost[i]` is fuel needed to travel from `i` to
`i+1` (wrapping). Return the starting index for a full clockwise loop, or `-1`
if impossible. A valid solution is guaranteed unique.

Why it's medium: the trick is recognizing that if total gas ≥ total cost, a
solution must exist, and the right start is the station right after the lowest
running deficit. Naive O(n²) tries every start; greedy is O(n).

Example:
  gas=[1,2,3,4,5], cost=[3,4,5,1,2] → 3
  gas=[2,3,4],     cost=[3,4,3]     → -1
