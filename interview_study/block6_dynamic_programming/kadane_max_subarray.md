PROBLEM: Max Subarray (Kadane)
===============================
Given an integer array `nums`, return the largest sum among all contiguous
subarrays. Empty input returns 0.

Why it's in this block: this is the canonical 1-D DP — state "best subarray
ending at i" with the recurrence `dp[i] = max(nums[i], dp[i-1] + nums[i])`,
then the textbook collapse from O(n) table to O(1) rolling scalar. The
all-negative case is the classic correctness trap.

Recurrence:
  dp[0] = nums[0]
  dp[i] = max(nums[i], dp[i-1] + nums[i])
  answer = max(dp[0..n-1])

Rolling form (drop the table, keep one scalar):
  curr = max(nums[i], curr + nums[i])
  best = max(best, curr)

Example:
  [-2,1,-3,4,-1,2,1,-5,4] -> 6   (subarray [4,-1,2,1])
  [1]                     -> 1
  [5,4,-1,7,8]            -> 23
  [-3,-1,-4,-2]           -> -1  (best single element when everything is negative)

See also: `block0_warmups/1_easy/max_subarray.md` for the warmup framing of the
same algorithm — this entry is the DP-recurrence view of that solution.
