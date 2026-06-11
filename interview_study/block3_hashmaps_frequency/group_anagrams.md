PROBLEM: Group Anagrams
========================
Given a slice of strings, group all anagrams together.
Return the groups in any order; elements within groups in any order.

Example:
  Input:  ["eat","tea","tan","ate","nat","bat"]
  Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Two approaches:
  1. Sort each word -> use sorted word as map key. O(n * k log k)
  2. Use [26]int frequency array as key. O(n * k)  <- preferred

The [26]int approach avoids string sorting entirely.
Go allows arrays as map keys directly since they are comparable.
