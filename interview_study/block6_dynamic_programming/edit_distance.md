PROBLEM: Edit Distance (Levenshtein)
====================================
Compute the minimum number of insertions, deletions, or substitutions to
convert string a into string b.

Why this block: extension of LCS to three operations. The base cases —
converting "" → b takes len(b) insertions, and a → "" takes len(a) deletions —
are easy to forget under pressure.

Recurrence (1-indexed):
  if a[i-1] == b[j-1]:
      dp[i][j] = dp[i-1][j-1]
  else:
      dp[i][j] = 1 + min(
          dp[i-1][j],     // delete from a
          dp[i][j-1],     // insert into a
          dp[i-1][j-1],   // substitute
      )

Examples:
  "horse", "ros"      → 3
  "intention", "execution" → 5
