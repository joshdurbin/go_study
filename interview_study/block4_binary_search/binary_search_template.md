PROBLEM: Search in Rotated Sorted Array
========================================
A sorted integer array was rotated at an unknown pivot.
Given the rotated array and a target, return the index or -1.

Example:
  Input:  nums=[4,5,6,7,0,1,2], target=0  -> 4
  Input:  nums=[4,5,6,7,0,1,2], target=3  -> -1

Key insight: one half of the array is always sorted.
  - Determine which half is sorted (compare nums[lo] to nums[mid])
  - Check if target falls in the sorted half
  - Eliminate the other half

Template:
  lo, hi := 0, len(nums)-1
  for lo <= hi {
      mid := lo + (hi-lo)/2
      ...
  }
  return -1

The mid overflow trick: lo + (hi-lo)/2 instead of (lo+hi)/2.
Always use this form.
