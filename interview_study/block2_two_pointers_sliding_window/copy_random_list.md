PROBLEM: Copy List with Random Pointer
======================================
Each node has Val, Next, and a Random pointer that points to any node in the list
or nil. Return a deep copy: every node duplicated, every Next and Random rewired
to clones (never to originals).

Why it's medium: the Random pointer breaks the natural single-pass order. The
hashmap solution (map[*Node]*Node, two passes) is straightforward; the
interleave-and-split trick achieves O(1) extra space and is the interview win.

Example:
  7 → 13 → 11 → 10 → 1
  randoms: nil, 0, 4, 2, 0
  Output: a fresh list with identical structure but all-new nodes.
