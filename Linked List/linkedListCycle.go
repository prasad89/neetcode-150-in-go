//go:build ignore

package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func buildWithCycle(arr []int, index int) *Node {
	var head, cycleNode *Node
	for i := len(arr) - 1; i >= 0; i-- {
		head = &Node{Value: arr[i], Next: head}
		if i == index {
			cycleNode = head
		}
	}
	tail := head
	for tail != nil && tail.Next != nil {
		tail = tail.Next
	}
	if index >= 0 && cycleNode != nil {
		tail.Next = cycleNode
	}
	return head
}

func hasCycle(head *Node) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	arr1, index1 := []int{1, 2, 3, 4}, 1
	head1 := buildWithCycle(arr1, index1)
	fmt.Println(hasCycle(head1))
	arr2, index2 := []int{1, 2}, -1
	head2 := buildWithCycle(arr2, index2)
	fmt.Println(hasCycle(head2))
}
