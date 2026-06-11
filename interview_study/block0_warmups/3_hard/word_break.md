PROBLEM: Word Break
===================
Given a string s and a dictionary of words, return true if s can be segmented
into a space-separated sequence of dictionary words.

Why it's hard: looks like backtracking but the naive recursion explodes —
needs memoization or 1-D DP. Tests recognizing DP from a non-numeric setup.

DP: dp[i] = true if s[:i] is breakable.
    dp[i] = OR over j of (dp[j] AND s[j:i] in dict)
    dp[0] = true (empty prefix).

Examples:
  s="leetcode",  dict=["leet","code"]         → true
  s="applepenapple", dict=["apple","pen"]     → true
  s="catsandog", dict=["cats","dog","sand","and","cat"] → false
