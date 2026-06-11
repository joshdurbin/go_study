//go:build ignore

package main

import "fmt"

// Multiple return values (idiomatic Go)
func minMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]
	for _, v := range nums[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

// Named return values (use sparingly — helps with godoc)
func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = fmt.Errorf("division by zero")
		return // naked return uses named values
	}
	result = a / b
	return
}

// Variadic function — receives args as a slice
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Functions are first-class values
func apply(f func(int) int, v int) int {
	return f(v)
}

// Closure captures variables from enclosing scope
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	min, max := minMax([]int{3, 1, 4, 1, 5, 9})
	fmt.Println(min, max) // 1 9

	result, err := divide(10, 3)
	fmt.Printf("%.4f %v\n", result, err) // 3.3333 <nil>

	_, err = divide(1, 0)
	fmt.Println(err) // division by zero

	fmt.Println(sum(1, 2, 3, 4, 5)) // 15
	// Spread a slice into variadic function
	nums := []int{1, 2, 3}
	fmt.Println(sum(nums...)) // 6

	double := func(x int) int { return x * 2 }
	fmt.Println(apply(double, 7)) // 14

	counter := makeCounter()
	fmt.Println(counter(), counter(), counter()) // 1 2 3
}
