//go:build ignore

package main

import "fmt"

func classify(n int) string {
	// if with init statement — x is scoped to the if block
	if x := n % 2; x == 0 {
		return "even"
	} else {
		_ = x // still accessible in else
		return "odd"
	}
}

func dayType(day string) string {
	switch day {
	case "Saturday", "Sunday":
		return "weekend"
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		return "weekday"
	default:
		return "unknown"
	}
}

func scoreGrade(score int) string {
	// switch with no expression == switch true (clean if-else chain)
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	default:
		return "F"
	}
}

func main() {
	fmt.Println(classify(4))   // even
	fmt.Println(classify(7))   // odd
	fmt.Println(dayType("Monday"))   // weekday
	fmt.Println(dayType("Sunday"))   // weekend
	fmt.Println(scoreGrade(85))      // B
	fmt.Println(scoreGrade(55))      // F
}
