//go:build ignore

package main

import "fmt"

// leastInterval: the most-frequent task dictates the schedule shape. Build
// (maxFreq-1) full frames of width (n+1), then add one slot per task tied at
// maxFreq. If tasks overflow the frame, length(tasks) is the answer.
// O(len(tasks)) time, O(1) space (26-letter histogram).
func leastInterval(tasks []byte, n int) int {
	var freq [26]int
	for _, t := range tasks {
		freq[t-'A']++
	}
	maxFreq, tiesAtMax := 0, 0
	for _, f := range freq {
		if f > maxFreq {
			maxFreq, tiesAtMax = f, 1
		} else if f == maxFreq {
			tiesAtMax++
		}
	}
	frame := (maxFreq-1)*(n+1) + tiesAtMax
	if len(tasks) > frame {
		return len(tasks)
	}
	return frame
}

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))                          // 8
	fmt.Println(leastInterval([]byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1))                          // 6
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2)) // 16
}
