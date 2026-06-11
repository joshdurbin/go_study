PROBLEM: Rate-Limited Worker Pool
=================================
Process a set of tasks with W parallel workers, but cap the *system-wide* rate
at R tasks per second. Worker count controls parallelism; the rate caps global
throughput so a downstream API isn't overwhelmed.

Why it's medium: combines worker pool with a token-bucket limiter. The key
insight is that the limiter must be shared across workers — a per-worker
ticker silently multiplies the rate by W.

Example:
  10 tasks, 2 workers, 2 tasks/sec → all 10 finish in ~5 seconds.
  W workers run in parallel but cannot collectively exceed R tasks/sec.
