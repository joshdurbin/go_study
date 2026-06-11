//go:build ignore

package main

import (
	"fmt"
	"sync"
)

// Concrete intro to four goroutine coordination patterns. Each is shown with
// the simplest possible setup so you can see the channel/sync.WaitGroup wiring.
// For the reusable, generic versions:
//   - Fan-out / fan-in over channels   → patterns/16_fan_in_fan_out.go
//   - Worker pool with bounded workers → patterns/02_worker_pool.go
//   - errgroup (first-error + cancel)  → patterns/13_errgroup.go
//   - Iterator and pipeline composition → patterns/10_iterator_and_pipeline.go

// fanOutSquares: concrete intro version — distribute squaring work across N
// workers and merge results onto one channel. See patterns/16 for the generic
// fanOut + fanIn pair that composes with any function and result type.
func fanOutSquares(jobs <-chan int, numWorkers int) <-chan int {
	results := make(chan int, numWorkers*2)
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				results <- j * j
			}
		}()
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	return results
}

// Pipeline: chain stages where output of one is input to the next
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// semaphore pattern: limit concurrent goroutines
func withSemaphore() {
	sem := make(chan struct{}, 3) // max 3 concurrent
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem <- struct{}{} // acquire
			defer func() { <-sem }() // release
			// only 3 goroutines are in this section at any time
			_ = id
		}(i)
	}
	wg.Wait()
	fmt.Println("semaphore done")
}

// Minimal errgroup-style runner: wait for all tasks, return the first non-nil
// error seen. The production-grade version (with sibling cancellation and a
// shared context) is in patterns/13_errgroup.go.
func runWithErrors(tasks []func() error) error {
	errs := make(chan error, len(tasks))
	var wg sync.WaitGroup
	for _, t := range tasks {
		wg.Add(1)
		t := t // capture loop variable
		go func() {
			defer wg.Done()
			errs <- t()
		}()
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// Fan-out
	jobs := make(chan int, 9)
	for i := 1; i <= 9; i++ {
		jobs <- i
	}
	close(jobs)
	results := fanOutSquares(jobs, 3)
	sum := 0
	for r := range results {
		sum += r
	}
	fmt.Println("fan-out sum of squares:", sum) // 285

	// Pipeline
	c := generate(2, 3, 4)
	out := square(c)
	for v := range out {
		fmt.Print(v, " ")
	}
	fmt.Println() // 4 9 16

	withSemaphore()

	// errgroup
	tasks := []func() error{
		func() error { return nil },
		func() error { return nil },
		func() error { return fmt.Errorf("task failed") },
	}
	if err := runWithErrors(tasks); err != nil {
		fmt.Println("error:", err)
	}
}
