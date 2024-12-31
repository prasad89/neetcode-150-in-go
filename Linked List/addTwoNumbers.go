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

func traverse(head *Node) []int {
	var result []int
	for current := head; current != nil; current = current.Next {
		result = append(result, current.Value)
	}
	return result
}

func addTwoNumbers(head1, head2 *Node) *Node {
	dummy := &Node{}
	curr, carry := dummy, 0
	for head1 != nil || head2 != nil || carry > 0 {
		sum := carry
		if head1 != nil {
			sum += head1.Value
			head1 = head1.Next
		}
		if head2 != nil {
			sum += head2.Value
			head2 = head2.Next
		}
		curr.Next = &Node{Value: sum % 10}
		curr = curr.Next
		carry = sum / 10
	}
	return dummy.Next
}

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	head1 := build(arr1)
	head2 := build(arr2)
	sumHead1 := addTwoNumbers(head1, head2)
	fmt.Println(traverse(sumHead1))
	arr3 := []int{9}
	arr4 := []int{9}
	head3 := build(arr3)
	head4 := build(arr4)
	sumHead2 := addTwoNumbers(head3, head4)
	fmt.Println(traverse(sumHead2))
}
