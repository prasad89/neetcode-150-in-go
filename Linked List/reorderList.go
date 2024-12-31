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

func reorderList(head *Node) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	secondHalf := slow.Next
	slow.Next = nil
	var prevSecondHalf *Node
	for secondHalf != nil {
		next := secondHalf.Next
		secondHalf.Next = prevSecondHalf
		prevSecondHalf = secondHalf
		secondHalf = next
	}
	first := head
	secondHalf = prevSecondHalf
	for secondHalf != nil {
		nextFirst, nextSecond := first.Next, secondHalf.Next
		first.Next = secondHalf
		secondHalf.Next = nextFirst
		first = nextFirst
		secondHalf = nextSecond
	}
}

func main() {
	arr1 := []int{2, 4, 6, 8}
	head1 := build(arr1)
	reorderList(head1)
	traverse(head1)
	arr2 := []int{2, 4, 6, 8, 10}
	head2 := build(arr2)
	reorderList(head2)
	traverse(head2)
}
