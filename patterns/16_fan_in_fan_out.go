//go:build ignore

package main

import (
	"fmt"
	"sync"
)

// FAN-OUT / FAN-IN
// ================
// Problem: you have a stream of work items and an expensive per-item operation.
// Process them concurrently across N workers, then merge results back into a
// single stream — in the original style of work scheduling, not as a worker
// pool with task queue.
//
// Fan-out: split one input channel into N concurrent processors.
// Fan-in:  merge those N output channels back into one.
//
// vs. worker_pool: worker pool fan-outs internally and you submit jobs.
// Fan-in/out is the lower-level channel-composition pattern that makes worker
// pools (and pipelines) possible.

// fanOut launches n goroutines, each reading from `in` and writing the result
// of fn to its own output channel. Returns all n output channels.
func fanOut[T, R any](in <-chan T, n int, fn func(T) R) []<-chan R {
	outs := make([]<-chan R, n)
	for i := 0; i < n; i++ {
		ch := make(chan R)
		outs[i] = ch
		go func() {
			defer close(ch)
			for v := range in {
				ch <- fn(v)
			}
		}()
	}
	return outs
}

// fanIn merges multiple channels into one. The merged channel closes once every
// input channel has closed.
func fanIn[R any](chans ...<-chan R) <-chan R {
	out := make(chan R)
	var wg sync.WaitGroup
	wg.Add(len(chans))
	for _, c := range chans {
		c := c
		go func() {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	// Producer: integers 1..10.
	in := make(chan int)
	go func() {
		defer close(in)
		for i := 1; i <= 10; i++ {
			in <- i
		}
	}()

	// Fan out across 4 workers that compute squares.
	workers := fanOut(in, 4, func(x int) int { return x * x })

	// Fan in to a single stream and collect.
	merged := fanIn(workers...)
	var results []int
	for v := range merged {
		results = append(results, v)
	}
	// Order is not guaranteed — workers process in parallel.
	fmt.Println("got", len(results), "squares:", results)
}
