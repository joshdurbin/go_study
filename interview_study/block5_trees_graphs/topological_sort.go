//go:build ignore

package main

import "fmt"

func canFinish(numCourses int, prerequisites [][2]int) bool {
	adj := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	for _, pre := range prerequisites {
		course, prereq := pre[0], pre[1]
		adj[prereq] = append(adj[prereq], course)
		inDegree[course]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	processed := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		processed++
		for _, neighbor := range adj[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return processed == numCourses
}

func main() {
	fmt.Println(canFinish(2, [][2]int{{1, 0}}))                  // true
	fmt.Println(canFinish(2, [][2]int{{1, 0}, {0, 1}}))          // false
	fmt.Println(canFinish(4, [][2]int{{1, 0}, {2, 1}, {3, 2}}))  // true
}
