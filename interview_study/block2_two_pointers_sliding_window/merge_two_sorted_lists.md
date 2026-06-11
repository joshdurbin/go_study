PROBLEM: Merge Two Sorted Lists
===============================
Given two sorted linked lists, merge them into one sorted list by splicing nodes.
Do not allocate new nodes.

Why this block: two-pointer pattern on two sequences instead of one. Dummy head
node simplifies the "first attach" edge case.

Example:
  1→2→4, 1→3→4 ⇒ 1→1→2→3→4→4
  [], 0        ⇒ 0
  [], []       ⇒ []
