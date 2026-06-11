PROBLEM: Generate Parentheses
=============================
Given n pairs of parentheses, return all combinations of well-formed
parentheses. Length of every result is 2n.

Why it's medium: classic backtracking with a constraint check. Generating all
2^(2n) strings then filtering is wasteful — the trick is to prune at each step
by tracking how many `(` and `)` you've placed. You can only close what you've
already opened.

Example:
  n=3 → ["((()))","(()())","(())()","()(())","()()()"]
  n=1 → ["()"]
