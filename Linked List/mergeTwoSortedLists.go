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

func mergeTwoLists(head1, head2 *Node) *Node {
	dummy := &Node{}
	current := dummy
	for head1 != nil && head2 != nil {
		if head1.Value < head2.Value {
			current.Next = head1
			head1 = head1.Next
		} else {
			current.Next = head2
			head2 = head2.Next
		}
		current = current.Next
	}
	if head1 != nil {
		current.Next = head1
	}
	if head2 != nil {
		current.Next = head2
	}
	return dummy.Next
}

func main() {
	list1 := []int{1, 2, 4}
	head1 := build(list1)
	list2 := []int{1, 3, 5}
	head2 := build(list2)
	mergedHead1 := mergeTwoLists(head1, head2)
	traverse(mergedHead1)
	list3 := []int{}
	head3 := build(list3)
	list4 := []int{1, 2}
	head4 := build(list4)
	mergedHead2 := mergeTwoLists(head3, head4)
	traverse(mergedHead2)
	list5 := []int{}
	head5 := build(list5)
	list6 := []int{}
	head6 := build(list6)
	mergedHead3 := mergeTwoLists(head5, head6)
	traverse(mergedHead3)
}
