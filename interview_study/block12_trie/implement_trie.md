PROBLEM: Implement Trie (Prefix Tree)
=====================================
Implement a Trie with Insert(word), Search(word) (exact match), and
StartsWith(prefix). Lowercase ASCII only.

Why it's medium: the data structure is unfamiliar to most candidates and
distinguishing "word ends here" (isEnd) from "prefix exists" is the trap. Every
op is O(L) — independent of the number of words stored.

Example:
  t := NewTrie()
  t.Insert("apple")
  t.Search("apple")      → true
  t.Search("app")        → false   // "app" was never inserted
  t.StartsWith("app")    → true    // "apple" has prefix "app"
  t.Insert("app")
  t.Search("app")        → true

Node shape:
  children [26]*Trie   // index by c - 'a'
  isEnd    bool        // true on the last char of an inserted word
