//go:build ignore

package main

import (
	"fmt"
)

// runPipeline launches numProducers goroutines (each pushing itemsPerProducer
// integers into a buffered channel of size bufSize) and numConsumers goroutines
// that read and count items. Returns the total processed.
//
// Constraints:
//   - producers own the channel; consumers must NOT close it
//   - no deadlocks, no leaked goroutines
//   - return only after every item is consumed
func runPipeline(numProducers, itemsPerProducer, numConsumers, bufSize int) int64 {
	// TODO: implement
	return 0
}

func main() {
	total := runPipeline(3, 4, 2, 2)
	fmt.Printf("processed %d items\n", total)
	// expect: processed 12 items
}
