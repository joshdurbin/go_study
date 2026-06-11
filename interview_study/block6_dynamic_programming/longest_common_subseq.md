PROBLEM: Longest Common Subsequence
===================================
Given two strings a and b, return the length of their longest common subsequence.
(A subsequence keeps relative order but does not need to be contiguous.)

Why this block: archetypal 2-D string DP. Drives the intuition for diff,
edit distance, and many bioinformatics-flavored problems.

Recurrence:
  dp[i][j] = dp[i-1][j-1] + 1            if a[i-1] == b[j-1]
           = max(dp[i-1][j], dp[i][j-1]) otherwise

Examples:
  "abcde", "ace" → 3 ("ace")
  "abc",   "abc" → 3
  "abc",   "def" → 0
