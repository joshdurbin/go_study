PROBLEM: Capacity To Ship Packages Within D Days
=================================================
Packages with given weights must be shipped in order each day.
Find the minimum ship capacity to ship all packages within D days.

Example:
  Input:  weights=[1,2,3,4,5,6,7,8,9,10], days=5  -> 15
  Input:  weights=[3,2,2,4,1,4], days=3            -> 6

Pattern: binary search on the answer space
  - lo = max(weights)  [must fit the heaviest single package]
  - hi = sum(weights)  [carry everything in one day]
  - Write a feasibility function: canShip(capacity) bool
  - Binary search: if canShip(mid), try smaller (hi=mid); else try larger (lo=mid+1)
  - Loop exits when lo == hi, which is the minimum valid capacity

This pattern transfers directly to: Koko eating bananas, minimum days to make bouquets,
split array largest sum, and many more.
