PROBLEM: Design Add and Search Words Data Structure
===================================================
Build a dictionary supporting AddWord(word) and Search(word). Search may
contain '.' wildcards, each matching any single letter.

Why it's medium: combines a standard trie with recursive DFS branching at
wildcard positions. The trick is recognizing that '.' forces you to try every
non-nil child — wildcards turn a linear lookup into a bounded fan-out.

Example:
  d := NewWordDictionary()
  d.AddWord("bad"); d.AddWord("dad"); d.AddWord("mad")
  d.Search("pad")  → false
  d.Search("bad")  → true
  d.Search(".ad")  → true
  d.Search("b..")  → true
  d.Search("b.z")  → false   // no inserted word matches
