PROBLEM: Word Search II
=======================
Given an m x n char board and a list of words, return every word that can be
formed by a path of 4-directionally adjacent cells (no cell reused per word).

Why it's hard: naive "run Word Search I once per word" is O(words * m*n*4^L) and
TLEs. The unlock is building a trie of the dictionary and carrying a trie
pointer through one DFS over the board — failed prefixes prune entire subtrees.
Also requires the in-place visited trick ('#' marker, restore on backtrack)
and de-duplicating results when the same word can be assembled multiple ways.

Example:
  board = [["o","a","a","n"],
           ["e","t","a","e"],
           ["i","h","k","r"],
           ["i","f","l","v"]]
  words = ["oath","pea","eat","rain"]
  → ["eat","oath"]

Key trick: store the full word string on the terminal trie node. When you hit a
node with non-empty word, append it to the result AND blank out the field so
duplicates aren't emitted.
