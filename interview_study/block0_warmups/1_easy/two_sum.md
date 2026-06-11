PROBLEM: Two Sum
================
Given an array of ints and a target, return the indices of the two numbers that
add up to target. Assume exactly one solution exists.

Why it's a warm-up: the canonical hashmap-lookup problem. Naive O(n²) two loops
vs. O(n) single-pass with a map of seen value → index.

Example:
  nums=[2,7,11,15], target=9 → [0,1]
  nums=[3,2,4],     target=6 → [1,2]
