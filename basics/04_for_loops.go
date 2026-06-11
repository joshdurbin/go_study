//go:build ignore

package main

import "fmt"

func main() {
	// Classic C-style for
	for i := 0; i < 3; i++ {
		fmt.Println("classic:", i)
	}

	// While-style (condition only)
	n := 1
	for n < 8 {
		n *= 2
	}
	fmt.Println("while-style result:", n)

	// Infinite loop with break
	count := 0
	for {
		count++
		if count == 3 {
			break
		}
	}
	fmt.Println("break at:", count)

	// continue skips to next iteration
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ") // 1 3 5
	}
	fmt.Println()

	// range over slice (index, value)
	nums := []int{10, 20, 30}
	for i, v := range nums {
		fmt.Printf("nums[%d] = %d\n", i, v)
	}

	// range over map (order not guaranteed)
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Printf("%s -> %d\n", k, v)
	}

	// range over string iterates runes (Unicode code points)
	for i, r := range "go!" {
		fmt.Printf("%d: %c\n", i, r)
	}
}
