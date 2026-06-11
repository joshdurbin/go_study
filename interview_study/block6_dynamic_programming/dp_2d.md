PROBLEM: Unique Paths + Longest Common Subsequence (2D DP)
===========================================================
PART A: Unique Paths
A robot starts at top-left of an m x n grid, can only move right or down.
How many unique paths to the bottom-right corner?

  Input:  m=3, n=7  -> 28
  Input:  m=3, n=2  -> 3

  dp[i][j] = dp[i-1][j] + dp[i][j-1]
  Base: first row and first column are all 1 (only one path to reach them).

PART B: Longest Common Subsequence
Given two strings, find the length of their longest common subsequence.
A subsequence is formed by deleting some characters without changing order.

  "abcde" and "ace"  -> 3 ("ace")
  "abc" and "def"    -> 0

  dp[i][j] = LCS length of s1[:i] and s2[:j]
  If s1[i-1] == s2[j-1]: dp[i][j] = dp[i-1][j-1] + 1
  Else:                   dp[i][j] = max(dp[i-1][j], dp[i][j-1])

Interview tip: always sketch the dp table on paper/whiteboard first.
Label rows with s1 chars, columns with s2 chars. Fill it in by hand for
a small example before writing code.
