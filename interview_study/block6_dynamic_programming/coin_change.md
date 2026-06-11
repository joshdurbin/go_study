PROBLEM: Coin Change (Minimum Coins)
====================================
Given coins of distinct denominations and a target amount, return the fewest
coins needed to make the amount. Return -1 if it cannot be made.

Why this block: classic unbounded-knapsack-style 1-D DP.
  dp[i] = min over coins c<=i of (dp[i-c] + 1)
  dp[0] = 0

Examples:
  coins=[1,2,5],   amount=11 → 3   (5+5+1)
  coins=[2],       amount=3  → -1
  coins=[1],       amount=0  → 0
