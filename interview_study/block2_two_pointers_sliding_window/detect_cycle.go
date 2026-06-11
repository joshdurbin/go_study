//go:build ignore

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// detectCycle returns the node where the cycle begins, or nil if none.
// Floyd's Tortoise & Hare: phase 1 finds a meeting point inside the cycle,
// phase 2 walks from head and from the meeting point at equal speed; they
// converge at the cycle's entry. O(n) time, O(1) space.
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

func build(vals []int, cycleAt int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	var cycleStart *ListNode
	for i, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
		if i == cycleAt {
			cycleStart = cur
		}
	}
	if cycleStart != nil {
		cur.Next = cycleStart
	}
	return dummy.Next
}

func nodeVal(n *ListNode) string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("%d", n.Val)
}

func main() {
	fmt.Println(nodeVal(detectCycle(build([]int{1, 2, 3, 4}, -1)))) // nil
	fmt.Println(nodeVal(detectCycle(build([]int{1, 2, 3, 4}, 1))))  // 2
	fmt.Println(nodeVal(detectCycle(build([]int{3, 2, 0, -4}, 1)))) // 2
}
