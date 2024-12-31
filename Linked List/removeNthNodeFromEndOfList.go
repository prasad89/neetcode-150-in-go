//go:build ignore

package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func build(arr []int) *Node {
	var head *Node
	for i := len(arr) - 1; i >= 0; i-- {
		head = &Node{Value: arr[i], Next: head}
	}
	return head
}

func traverse(head *Node) {
	if head == nil {
		fmt.Println("The list is empty.")
		return
	}
	current := head
	for current != nil {
		if current.Next != nil {
			fmt.Print(current.Value, "->")
		} else {
			fmt.Print(current.Value)
		}
		current = current.Next
	}
	fmt.Println()
}

func removeNthFromEnd(head *Node, n int) *Node {
    slow, fast := head, head
    for i := 0; i < n; i++ {
        fast = fast.Next
    }
    if fast == nil {
        return head.Next
    }
    for fast.Next != nil {
        slow = slow.Next
        fast = fast.Next
    }
    slow.Next = slow.Next.Next
    return head
}

func main() {
	arr1, n1 := []int{1, 2, 3, 4}, 2
	head1 := build(arr1)
	removed1 := removeNthFromEnd(head1, n1)
	traverse(removed1)
	arr2, n2 := []int{5}, 1
	head2 := build(arr2)
	removed2 := removeNthFromEnd(head2, n2)
	traverse(removed2)
	arr3, n3 := []int{1, 2}, 2
	head3 := build(arr3)
	removed3 := removeNthFromEnd(head3, n3)
	traverse(removed3)
}
