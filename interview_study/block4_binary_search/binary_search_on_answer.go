//go:build ignore

package main

import "fmt"

func canShip(weights []int, capacity, maxDays int) bool {
	days := 1
	load := 0
	for _, w := range weights {
		if load+w > capacity {
			days++
			load = 0
		}
		load += w
	}
	return days <= maxDays
}

func shipWithinDays(weights []int, days int) int {
	lo, hi := 0, 0
	for _, w := range weights {
		if w > lo {
			lo = w
		}
		hi += w
	}
	for lo < hi {
		mid := lo + (hi-lo)/2
		if canShip(weights, mid, days) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)) // 15
	fmt.Println(shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))               // 6
	fmt.Println(shipWithinDays([]int{1, 2, 3, 1, 1}, 4))                  // 3
}
