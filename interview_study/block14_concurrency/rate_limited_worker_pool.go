//go:build ignore

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// runRateLimitedPool: W workers process tasks at a global rate of ratePerSec.
// A single ticker drips tokens into a shared channel; each worker MUST consume
// one token per task. This caps the system-wide throughput regardless of W.
//
// Why this design: a per-worker limiter would let total = W * rate. The shared
// token channel decouples worker count (parallelism) from rate (throughput).
func runRateLimitedPool(tasks []int, workers int, ratePerSec int) (processed int64, elapsed time.Duration) {
	jobs := make(chan int, len(tasks))
	tokens := make(chan struct{}, ratePerSec) // small burst = one second's worth

	// Token source: emit one token every 1/ratePerSec.
	stop := make(chan struct{})
	tick := time.NewTicker(time.Second / time.Duration(ratePerSec))
	go func() {
		defer tick.Stop()
		for {
			select {
			case <-tick.C:
				select { // non-blocking: if bucket full, drop
				case tokens <- struct{}{}:
				default:
				}
			case <-stop:
				return
			}
		}
	}()

	// Seed the bucket with one token so the first task fires immediately.
	tokens <- struct{}{}

	start := time.Now()

	// Workers: pull job, wait for a token, "do work".
	var wg sync.WaitGroup
	var count int64
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range jobs {
				<-tokens // throttles the whole pool
				atomic.AddInt64(&count, 1)
			}
		}()
	}

	for _, t := range tasks {
		jobs <- t
	}
	close(jobs)

	wg.Wait()
	close(stop)
	return atomic.LoadInt64(&count), time.Since(start)
}

func main() {
	tasks := make([]int, 10)
	for i := range tasks {
		tasks[i] = i
	}

	processed, elapsed := runRateLimitedPool(tasks, 2, 2)

	// 10 tasks at 2/sec with 1 seeded token ≈ 1 + 9/2 ≈ ~4.5s. Allow 3.5–6s.
	withinExpected := elapsed >= 3500*time.Millisecond && elapsed <= 6500*time.Millisecond

	fmt.Printf("processed=%d in_window=%v\n", processed, withinExpected)
	// expected:
	// processed=10 in_window=true
}
