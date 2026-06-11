//go:build ignore

package main

import (
	"fmt"
	"time"
)

// runRateLimitedPool processes tasks with `workers` parallel goroutines, but
// the entire pool may complete at most ratePerSec tasks per second.
//
// Constraints:
//   - workers controls parallelism; ratePerSec controls global throughput
//   - the token source goroutine must not leak after the work finishes
//   - return total processed and total elapsed time
func runRateLimitedPool(tasks []int, workers int, ratePerSec int) (processed int64, elapsed time.Duration) {
	// TODO: implement using a single ticker + shared token channel + worker pool
	return 0, 0
}

func main() {
	tasks := make([]int, 10)
	for i := range tasks {
		tasks[i] = i
	}

	processed, elapsed := runRateLimitedPool(tasks, 2, 2)

	// 10 tasks at 2/sec with 1 seeded token ≈ ~4.5s. Allow 3.5–6s.
	withinExpected := elapsed >= 3500*time.Millisecond && elapsed <= 6500*time.Millisecond

	fmt.Printf("processed=%d in_window=%v\n", processed, withinExpected)
	// expect: processed=10 in_window=true
}
