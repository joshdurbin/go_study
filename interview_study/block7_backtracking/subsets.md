PROBLEM: Subsets (Power Set)
============================
Given an array of distinct integers `nums`, return all possible subsets (the
power set). The solution set must not contain duplicate subsets.

Why it's medium: the canonical backtracking template. Teaches the "choose / un-choose"
pattern, the importance of copying the path before appending, and the `start` index
trick that prevents permutations from sneaking in.

Example:
  nums=[1,2,3] → [[],[1],[1,2],[1,2,3],[1,3],[2],[2,3],[3]]
  nums=[0]     → [[],[0]]
