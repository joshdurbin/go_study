PROBLEM: Longest Substring Without Repeating Characters
========================================================
Given a string, find the length of the longest substring that contains
no repeating characters.

Example:
  Input:  "abcabcbb"  -> 3 ("abc")
  Input:  "bbbbb"     -> 1 ("b")
  Input:  "pwwkew"    -> 3 ("wke")

Pattern: variable-size sliding window with a map
  - charIndex map tracks most recent index of each character
  - left pointer jumps past the last seen duplicate
  - Never move left backwards (use max)

Time: O(n)  Space: O(min(n, alphabet))
