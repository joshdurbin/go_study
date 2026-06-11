//go:build ignore

package main

import "fmt"

// canAttendMeetings reports whether a person could attend every meeting,
// i.e. no two intervals overlap. Touching (end == next.start) is allowed.
func canAttendMeetings(intervals [][]int) bool {
	// TODO: implement
	return false
}

func main() {
	fmt.Println(canAttendMeetings([][]int{{0, 30}, {5, 10}, {15, 20}})) // expect false
	fmt.Println(canAttendMeetings([][]int{{7, 10}, {2, 4}}))            // expect true
	fmt.Println(canAttendMeetings([][]int{{1, 5}, {5, 8}}))             // expect true
}
