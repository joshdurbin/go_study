PROBLEM: Rotate Array
=====================
Rotate an array to the right by k steps, in-place, with O(1) extra space.

Why it's medium: the elegant trick — reverse the whole array, then reverse
the first k and the rest — is non-obvious and worth memorizing.

Example:
  [1,2,3,4,5,6,7], k=3 → [5,6,7,1,2,3,4]
  [-1,-100,3,99],  k=2 → [3,99,-1,-100]
