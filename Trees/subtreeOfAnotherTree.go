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

func isSubtree(root, subRoot *Node) bool {
	if root == nil || subRoot == nil {
		return false
	}
	return isSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(root1, root2 *Node) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil || root1.Value != root2.Value {
		return false
	}
	return isSameTree(root1.Left, root2.Left) && isSameTree(root1.Right, root2.Right)
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{2, 4, 5}
	root1 := build(arr1)
	root2 := build(arr2)
	fmt.Println(isSubtree(root1, root2))
	arr3 := []int{1, 2, 3, 4, 5, -1, -1, 6}
	arr4 := []int{2, 4, 5}
	root3 := build(arr3)
	root4 := build(arr4)
	fmt.Println(isSubtree(root3, root4))
}
