//go:build ignore

package main

import "fmt"

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

func reverseList(head *Node) *Node {
	var prev *Node
	curr := head
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return prev
}

func main() {
	arr1 := []int{0, 1, 2, 3}
	head1 := build(arr1)
	reversedHead1 := reverseList(head1)
	traverse(reversedHead1)
	arr2 := []int{}
	head2 := build(arr2)
	reversedHead2 := reverseList(head2)
	traverse(reversedHead2)
}
