PROBLEM: Binary Tree — Max Depth & Path Sum
============================================
Part A: Find the maximum depth (nodes along longest root-to-leaf path).

Part B: Given a target sum, determine if there's a root-to-leaf path
        where all node values sum to target.

Example tree:
        5
       / \
      4   8
     /   / \
    11  13   4
   /  \       \
  7    2       1

  Max depth: 4
  hasPathSum(tree, 22): true  (5->4->11->2)
  hasPathSum(tree, 26): false

Practice both recursive (clean) and iterative (explicit stack) DFS.
Knowing both matters — interviewers often ask for the iterative form
after you nail the recursive one.
