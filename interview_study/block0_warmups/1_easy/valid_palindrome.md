PROBLEM: Valid Palindrome
=========================
Return true if a string is a palindrome considering only alphanumeric characters
and ignoring case.

Why it's a warm-up: opposite-ends two-pointer pattern, plus rune handling and
unicode-aware case folding.

Example:
  "A man, a plan, a canal: Panama" → true
  "race a car" → false
  "" → true
