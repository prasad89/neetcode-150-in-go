//go:build ignore

package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func build(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	root := &Node{Value: arr[0]}
	queue := []*Node{root}
	ptr := 1
	for len(queue) > 0 && ptr < len(arr) {
		curr := queue[0]
		queue = queue[1:]
		if ptr < len(arr) && arr[ptr] != -1 {
			curr.Left = &Node{Value: arr[ptr]}
			queue = append(queue, curr.Left)
		}
		ptr++
		if ptr < len(arr) && arr[ptr] != -1 {
			curr.Right = &Node{Value: arr[ptr]}
			queue = append(queue, curr.Right)
		}
		ptr++
	}
	return root
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	leftHeight := maxDepth(root.Left)
	rightHeight := maxDepth(root.Right)
	return 1 + max(leftHeight, rightHeight)
}

func main() {
	arr1 := []int{1, 2, 3, -1, -1, 4}
	root1 := build(arr1)
	fmt.Println(maxDepth(root1))
	arr2 := []int{}
	root2 := build(arr2)
	fmt.Println(maxDepth(root2))
}
