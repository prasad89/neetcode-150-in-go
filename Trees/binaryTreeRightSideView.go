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
	for ptr < len(arr) {
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

func rightSideView(root *Node) []int {
	if root == nil {
		return nil
	}
	result := []int{}
	queue := []*Node{root}
	for len(queue) > 0 {
		level := []int{}
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			curr := queue[0]
			queue = queue[1:]
			level = append(level, curr.Value)
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		result = append(result, level[len(level)-1])
	}
	return result
}

func main() {
	arr1 := []int{1, 2, 3}
	root1 := build(arr1)
	fmt.Println(rightSideView(root1))
	arr2 := []int{1, 2, 3, 4, 5, 6, 7}
	root2 := build(arr2)
	fmt.Println(rightSideView(root2))
}
