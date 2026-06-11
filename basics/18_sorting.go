//go:build ignore

package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Sorting built-in types
	ints := []int{5, 2, 8, 1, 9, 3}
	sort.Ints(ints)
	fmt.Println(ints) // [1 2 3 5 8 9]

	strs := []string{"banana", "apple", "cherry"}
	sort.Strings(strs)
	fmt.Println(strs) // [apple banana cherry]

	// Custom sort with sort.Slice
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Carol", 35},
		{"Dave", 25},
	}
	// Sort by age ascending, then name ascending for ties
	sort.Slice(people, func(i, j int) bool {
		if people[i].Age != people[j].Age {
			return people[i].Age < people[j].Age
		}
		return people[i].Name < people[j].Name
	})
	for _, p := range people {
		fmt.Printf("%s %d\n", p.Name, p.Age)
	}

	// sort.Search: binary search on a sorted slice
	// Returns smallest index i in [0,n) where f(i) is true
	nums := []int{1, 3, 5, 7, 9, 11}
	target := 7
	idx := sort.SearchInts(nums, target)
	fmt.Printf("found %d at index %d\n", target, idx) // found 7 at index 3

	// Check if slice is sorted
	fmt.Println(sort.IntsAreSorted(ints)) // true

	// Reverse sort
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println(ints) // [9 8 5 3 2 1]
}
