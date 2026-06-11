//go:build ignore

package main

import "fmt"

func readFile(name string) (err error) {
	fmt.Printf("opening %s\n", name)
	// defer runs when the surrounding function returns, in LIFO order
	// Useful for cleanup — runs even if the function panics
	defer fmt.Printf("closing %s\n", name)

	if name == "" {
		return fmt.Errorf("empty filename")
	}
	fmt.Printf("reading %s\n", name)
	return nil
}

func deferOrder() {
	// Multiple defers run LIFO (last-in, first-out)
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	fmt.Println("function body")
}

func riskyOperation(crash bool) (result string) {
	defer func() {
		if r := recover(); r != nil {
			// recover() stops the panic and returns the panic value
			result = fmt.Sprintf("recovered: %v", r)
		}
	}()
	if crash {
		panic("something went very wrong")
	}
	return "success"
}

func main() {
	readFile("data.txt")
	fmt.Println("---")
	readFile("")
	fmt.Println("---")

	deferOrder()
	fmt.Println("---")

	fmt.Println(riskyOperation(false)) // success
	fmt.Println(riskyOperation(true))  // recovered: something went very wrong
}
