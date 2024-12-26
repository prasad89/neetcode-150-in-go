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

func maxPathSum(root *Node) int {
	res := 0
	var maxSum func(root *Node) int
	maxSum = func(root *Node) int {
		if root == nil {
			return 0
		}
		leftMax := maxSum(root.Left)
		rightMax := maxSum(root.Right)
		leftMax = max(leftMax, 0)
		rightMax = max(rightMax, 0)
		res = max(res, root.Value+leftMax+rightMax)
		return root.Value + max(leftMax, rightMax)
	}
	maxSum(root)
	return res
}

func main() {
	arr1 := []int{1, 2, 3}
	root1 := build(arr1)
	fmt.Println(maxPathSum(root1))
	arr2 := []int{-15, 10, 20, -1, -1, 15, 5, -5}
	root2 := build(arr2)
	fmt.Println(maxPathSum(root2))
}
