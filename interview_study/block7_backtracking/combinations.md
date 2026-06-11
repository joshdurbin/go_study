PROBLEM: Combinations
=====================
Given two integers `n` and `k`, return all possible k-length combinations of
the numbers in the range [1, n].

Why it's medium: a constrained version of subsets — same backtracking template
but only record at depth k. Good chance to demonstrate pruning the loop bound
so you don't recurse into branches that can't fill the path.

Example:
  n=4, k=2 → [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
  n=3, k=1 → [[1],[2],[3]]
