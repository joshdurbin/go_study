PROBLEM: Binary Tree Level Order Traversal
===========================================
Given a binary tree root, return level order traversal as [][]int.

Example:
  Tree:  3
        / \
       9  20
         /  \
        15   7

  Output: [[3],[9,20],[15,7]]

Pattern: BFS with a queue (slice-as-queue in Go).
  - Snapshot queue length at the start of each iteration = nodes at this level
  - Process exactly that many nodes, enqueue their children
  - Each pass = one level

Also implement: right side view (last node visible at each level from the right).

Time: O(n)  Space: O(n)
