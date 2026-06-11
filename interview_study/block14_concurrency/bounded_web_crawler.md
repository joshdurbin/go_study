PROBLEM: Bounded Web Crawler
============================
Given a list of URLs and a concurrency limit N, fetch them all with at most N
in flight at any moment. Return the bodies. The work is I/O-bound (mocked here
with a sleep).

Why it's medium: tests the "semaphore channel" idiom — a buffered channel of
`struct{}` as a counting semaphore. Also tests safe result collection without
data races (slice with goroutine-owned indices, or a results channel).

Example:
  10 urls, concurrency=3 → all 10 fetched, never more than 3 goroutines
  actively running, total time ≈ ceil(10/3) * per-call latency.
