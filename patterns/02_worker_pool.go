//go:build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

// WORKER POOL PATTERN
// ===================
// Problem: You have N tasks and want to process them with bounded concurrency —
// not one goroutine per task (could spawn thousands) and not one at a time.
//
// This is one of the most common Go concurrency patterns in real systems.
// Your Redis upgrade tooling, DB failover pipelines, GCP inventory scripts —
// all of these benefit from a worker pool.
//
// Core idea:
//   - Fixed pool of worker goroutines, all reading from a shared jobs channel
//   - Main goroutine sends jobs, closes the channel when done
//   - Workers send results to a results channel
//   - A separate goroutine waits for all workers then closes results
//
// Key: the jobs channel acts as the coordination primitive. No mutexes needed.

type Job struct {
	ID    int
	Input string
}

type Result struct {
	JobID  int
	Output string
	Err    error
}

// process simulates work — replace with your actual logic
func process(j Job) Result {
	time.Sleep(10 * time.Millisecond) // simulate I/O
	return Result{
		JobID:  j.ID,
		Output: fmt.Sprintf("processed: %s", j.Input),
	}
}

// WorkerPool runs numWorkers goroutines, each pulling from jobs and sending to results.
// Caller must close(jobs) after all jobs are sent.
// Results channel is closed automatically when all workers finish.
func WorkerPool(numWorkers int, jobs <-chan Job) <-chan Result {
	results := make(chan Result, numWorkers*2) // buffered to reduce blocking
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs { // range blocks until jobs is closed
				results <- process(job)
			}
		}(i)
	}

	// Closer goroutine: waits for all workers, then signals no more results
	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

// VARIANT: Worker pool with context for cancellation
// (in production you'd add context.Context to process() and check ctx.Done())

// VARIANT: Dynamic pool — adjust worker count at runtime
// Use a semaphore channel instead of fixed goroutines:
//   sem := make(chan struct{}, maxWorkers)
//   for _, job := range jobs {
//       sem <- struct{}{}
//       go func(j Job) { defer func() { <-sem }(); process(j) }(job)
//   }
// Simpler but slightly different semantics — goroutines are created per job.

func main() {
	jobs := make(chan Job, 20)

	// Send jobs — in real code this might be a DB query result set, file list, etc.
	go func() {
		for i := 0; i < 20; i++ {
			jobs <- Job{ID: i, Input: fmt.Sprintf("item-%d", i)}
		}
		close(jobs) // CRITICAL: workers range over jobs; must close to unblock them
	}()

	start := time.Now()
	results := WorkerPool(5, jobs) // 5 workers, 20 jobs, 10ms each => ~40ms not 200ms

	// Collect results as they arrive
	var allResults []Result
	for r := range results { // range blocks until results is closed
		allResults = append(allResults, r)
	}

	fmt.Printf("processed %d jobs in %v with 5 workers\n", len(allResults), time.Since(start).Round(time.Millisecond))
	fmt.Println("sample:", allResults[0].Output)

	// GOTCHA: forgetting to close(jobs) causes workers to hang forever.
	// GOTCHA: not draining results before calling something that needs completion
	//         causes the closer goroutine to block on results <- process(job).
}
