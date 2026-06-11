PROBLEM: Valid Parentheses
==========================
Given a string containing only '(', ')', '{', '}', '[', ']', determine if the
brackets are balanced and correctly nested.

Why it's medium: classic stack problem. Tests pattern recognition — any
"matching" or "innermost first" requirement → stack.

Examples:
  "()"       → true
  "()[]{}"   → true
  "(]"       → false
  "([)]"     → false
  "{[]}"     → true
