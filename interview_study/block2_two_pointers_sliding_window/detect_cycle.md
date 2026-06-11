PROBLEM: Linked List Cycle II — Find Cycle Start
=================================================
Given the head of a linked list, return the node where the cycle begins, or nil if
there is no cycle. Solve in O(1) extra space.

Why it's medium: Floyd's two-phase trick is non-obvious. Phase 1 (detect) is easy;
phase 2 (locate entry) requires the algebraic insight that distance(head→entry)
equals distance(meeting→entry) when walking around the cycle.

Example:
  1 → 2 → 3 → 4, no cycle    ⇒ nil
  1 → 2 → 3 → 4 → (back to 2) ⇒ node 2
  3 → 2 → 0 → -4 → (back to 2)⇒ node 2
