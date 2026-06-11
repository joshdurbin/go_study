PROBLEM: Coin Change (1D DP — bottom-up tabulation)
=====================================================
Given coin denominations and an amount, return the fewest coins to make
that amount. Return -1 if impossible.

Example:
  coins=[1,2,5], amount=11  -> 3  (5+5+1)
  coins=[2], amount=3       -> -1
  coins=[1], amount=0       -> 0

Pattern: define dp[i] = minimum coins to make amount i
  dp[0] = 0  (zero coins for amount 0)
  dp[i] = min over all coins: dp[i-coin] + 1  (if coin <= i)
  Initialize dp[1..amount] = amount+1 as a sentinel for "impossible"

Also implement: Climbing Stairs
  dp[i] = number of ways to reach step i (taking 1 or 2 steps at a time)
  dp[i] = dp[i-1] + dp[i-2]
  Optimize to O(1) space by keeping only last two values.
