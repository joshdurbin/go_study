PROBLEM: Longest Substring Without Repeating Characters
=======================================================
Return the length of the longest substring of s containing no repeated characters.

Why it's medium: prime variable-size sliding window exercise. The trick is what
to track in the window — a map of char → last-seen-index lets you jump left
forward instead of advancing one step at a time.

Examples:
  "abcabcbb" → 3  ("abc")
  "bbbbb"    → 1  ("b")
  "pwwkew"   → 3  ("wke")
  ""         → 0
