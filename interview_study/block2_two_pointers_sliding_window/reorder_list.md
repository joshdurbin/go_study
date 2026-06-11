PROBLEM: Reorder List
=====================
Given L0→L1→...→Ln-1→Ln, reorder it in-place to L0→Ln→L1→Ln-1→L2→Ln-2→...
Modify pointers only — do not allocate new nodes.

Why it's medium: composes three classic linked-list moves — slow/fast midpoint,
in-place reverse, and two-pointer splice. Each is easy alone; chaining them with
correct boundary handling is the test.

Example:
  1 → 2 → 3 → 4       ⇒ 1 → 4 → 2 → 3
  1 → 2 → 3 → 4 → 5   ⇒ 1 → 5 → 2 → 4 → 3
  1                   ⇒ 1
