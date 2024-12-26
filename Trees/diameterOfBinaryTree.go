//go:build ignore

package main

import (
	"fmt"
)

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

func diameterOfBinaryTree(root *Node) int {
	res := 0
	var diameter func(root *Node) int
	diameter = func(root *Node) int {
		if root == nil {
			return 0
		}
		leftHeight := diameter(root.Left)
		rightHeight := diameter(root.Right)
		res = max(res, leftHeight+rightHeight)
		return 1 + max(leftHeight, rightHeight)
	}
	diameter(root)
	return res
}

func main() {
	arr1 := []int{1, -1, 2, 3, 4, 5}
	root1 := build(arr1)
	fmt.Println(diameterOfBinaryTree(root1))
	arr2 := []int{1, 2, 3}
	root2 := build(arr2)
	fmt.Println(diameterOfBinaryTree(root2))
}
