PROBLEM: Reverse a String
=========================
Reverse a UTF-8 string. Naive byte-reversal breaks multi-byte runes.

Why it's a warm-up: forces you to remember that Go strings are bytes, but ranging
over a string yields runes. Convert to []rune, swap ends, return string([]rune).

Example:
  Input:  "héllo"
  Output: "olléh"

Bonus: do it without allocating a second slice (in-place on []rune).
