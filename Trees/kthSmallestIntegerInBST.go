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

func kthSmallest(root *Node, k int) int {
	result, count := 0, k
	var smallest func(root *Node)
	smallest = func(root *Node) {
		if root == nil {
			return
		}
		smallest(root.Left)
		count--
		if count == 0 {
			result = root.Value
			return
		}
		smallest(root.Right)
	}
	smallest(root)
	return result
}

func main() {
	arr1, k1 := []int{2, 1, 3}, 1
	root1 := build(arr1)
	fmt.Println(kthSmallest(root1, k1))
	arr2, k2 := []int{4, 3, 5, 2, -1}, 4
	root2 := build(arr2)
	fmt.Println(kthSmallest(root2, k2))
}
