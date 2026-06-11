PROBLEM: Task Scheduler
=======================
Given `tasks []byte` (uppercase letters) and cooldown `n`, each CPU interval
runs one task or idles. The same task cannot run twice within `n` intervals.
Return the minimum total intervals to finish all tasks.

Why it's medium: looks like simulation with a heap (O(N log 26)), but the
closed-form answer is O(N). Place the most frequent task into `(maxFreq-1)`
frames of size `(n+1)`, append the ties, and floor at `len(tasks)`.

Example:
  tasks="AAABBB", n=2 → 8   (A B _ A B _ A B)
  tasks="ACABDB", n=1 → 6   (no idles needed)
