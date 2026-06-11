PROBLEM: Lowest Common Ancestor (Binary Tree)
=============================================
Given a binary tree and two distinct nodes p and q (guaranteed in the tree),
return their lowest common ancestor.

Why this block: classic post-order DFS. Each subtree reports whether it
"contains" p or q; the first node where both sides report true is the LCA.

Example tree:
        3
       / \
      5   1
     /|   |\
    6 2   0 8
      |\
      7 4

  LCA(5, 1) → 3
  LCA(5, 4) → 5  (a node can be its own descendant)
