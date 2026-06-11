//go:build ignore

package main

import "fmt"

// fizzbuzz returns the FizzBuzz sequence for 1..n as a slice of strings.
// Multiples of 3 → "Fizz", multiples of 5 → "Buzz", multiples of both → "FizzBuzz".
// All other numbers → the decimal representation of the number itself.
func fizzbuzz(n int) []string {
	// TODO: implement
	return nil
}

func main() {
	got := fizzbuzz(15)
	want := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8",
		"Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	fmt.Println("got: ", got)
	fmt.Println("want:", want)
}
