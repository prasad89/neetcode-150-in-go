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

func reverseKGroup(head *Node, k int) *Node {
	dummy := &Node{Next: head}
	groupPrev := dummy
	for {
		kth := getKthNode(groupPrev, k)
		if kth == nil {
			break
		}
		groupNext := kth.Next
		prev, curr := groupNext, groupPrev.Next
		for curr != groupNext {
			tmp := curr.Next
			curr.Next = prev
			prev = curr
			curr = tmp
		}
		tmp := groupPrev.Next
		groupPrev.Next = kth
		groupPrev = tmp

	}
	return dummy.Next
}

func getKthNode(curr *Node, k int) *Node {
	for curr != nil && k > 0 {
		curr = curr.Next
		k--
	}
	return curr
}

func main() {
	arr1, k1 := []int{1, 2, 3, 4, 5, 6}, 3
	head1 := build(arr1)
	reversedHead1 := reverseKGroup(head1, k1)
	traverse(reversedHead1)
	arr2, k2 := []int{1, 2, 3, 4, 5}, 3
	head2 := build(arr2)
	reversedHead2 := reverseKGroup(head2, k2)
	traverse(reversedHead2)
}
