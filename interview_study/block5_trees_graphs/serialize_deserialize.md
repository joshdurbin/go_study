PROBLEM: Serialize and Deserialize Binary Tree
==============================================
Design Serialize(root) → string and Deserialize(string) → root such that
deserialize(serialize(t)) reproduces t.

Why this block: combines preorder DFS with sentinel handling for nil children.
The "#" / "null" marker for empty children makes the structure fully recoverable
without needing inorder + preorder.

Example:
  Tree:
        1
       / \
      2   3
         / \
        4   5

  Serialize → "1,2,#,#,3,4,#,#,5,#,#"
