//go:build ignore

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// runPipeline: N producers fan items into one bounded channel; M consumers drain it.
// WaitGroup tracks producers so we know exactly when to close(ch); consumers exit on close.
// Total work O(N*items); peak memory bounded by channel capacity B.
func runPipeline(numProducers, itemsPerProducer, numConsumers, bufSize int) int64 {
	ch := make(chan int, bufSize)
	var producers sync.WaitGroup
	var consumers sync.WaitGroup
	var processed int64

	// Producers: each pushes itemsPerProducer integers, then signals done.
	for p := 0; p < numProducers; p++ {
		producers.Add(1)
		go func(id int) {
			defer producers.Done()
			for i := 0; i < itemsPerProducer; i++ {
				ch <- id*itemsPerProducer + i // back-pressure: blocks when buffer full
			}
		}(p)
	}

	// Closer: once all producers finished, close the channel so consumers can exit.
	go func() {
		producers.Wait()
		close(ch)
	}()

	// Consumers: range exits when ch is closed AND drained.
	for c := 0; c < numConsumers; c++ {
		consumers.Add(1)
		go func() {
			defer consumers.Done()
			for range ch {
				atomic.AddInt64(&processed, 1)
			}
		}()
	}

	consumers.Wait()
	return processed
}

func main() {
	// 3 producers x 4 items = 12 items; 2 consumers; channel buffer 2.
	total := runPipeline(3, 4, 2, 2)
	fmt.Printf("processed %d items\n", total)
	// expected:
	// processed 12 items
}
