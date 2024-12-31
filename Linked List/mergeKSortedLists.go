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

func mergeKLists(lists []*Node) *Node {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	left := mergeKLists(lists[:n/2])
	right := mergeKLists(lists[n/2:])
	return mergeTwoLists(left, right)
}

func mergeTwoLists(list1, list2 *Node) *Node {
	dummy := &Node{}
	curr := dummy
	for list1 != nil && list2 != nil {
		if list1.Value < list2.Value {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}
	if list1 != nil {
		curr.Next = list1
	}
	if list2 != nil {
		curr.Next = list2
	}
	return dummy.Next
}

func main() {
	lists1 := []*Node{
		build([]int{1, 2, 4}),
		build([]int{1, 3, 5}),
		build([]int{3, 6}),
	}
	mergedHead1 := mergeKLists(lists1)
	traverse(mergedHead1)
	lists2 := []*Node{
		build([]int{}),
		build([]int{}),
	}
	mergedHead2 := mergeKLists(lists2)
	traverse(mergedHead2)
	lists3 := []*Node{
		build([]int{}),
	}
	mergedHead3 := mergeKLists(lists3)
	traverse(mergedHead3)
}
