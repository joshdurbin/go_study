PROBLEM: 0/1 Knapsack
=====================
Given items with weights w[i] and values v[i], and a capacity W, maximize the
total value such that the total weight is ≤ W. Each item can be used at most once.

Why this block: foundational 2-D DP. The "0/1" twist (vs. unbounded knapsack)
is what makes the recurrence iterate items in the OUTER loop. With a 1-D
rolling array, iterate weight RIGHT TO LEFT to avoid re-using the same item.

Recurrence:
  dp[i][w] = max(
    dp[i-1][w],                            // skip item i
    dp[i-1][w - weights[i]] + values[i]    // take it (if it fits)
  )

Example:
  weights=[1,3,4,5], values=[1,4,5,7], capacity=7 → 9  (items 1 and 2: w=3+4=7, v=4+5=9)
