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

func goodNodes(root *Node) int {
	res := 0
	var count func(root *Node, maxValue int)
	count = func(root *Node, maxValue int) {
		if root == nil {
			return
		}
		if root.Value >= maxValue {
			res++
			maxValue = root.Value
		}
		count(root.Left, maxValue)
		count(root.Right, maxValue)
	}
	count(root, root.Value)
	return res
}

func main() {
	arr1 := []int{2, 1, 1, 3, -1, 1, 5}
	root1 := build(arr1)
	fmt.Println(goodNodes(root1))
	arr2 := []int{1, 2, 0, 3, 4}
	root2 := build(arr2)
	fmt.Println(goodNodes(root2))
}
