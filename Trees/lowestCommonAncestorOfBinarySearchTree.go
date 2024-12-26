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

func lowestCommonAncestor(root, p, q *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Value < p.Value && root.Value < q.Value {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if root.Value > p.Value && root.Value > q.Value {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}

func main() {
	arr1 := []int{5, 3, 8, 1, 4, 7, 9, -1, 2}
	root1 := build(arr1)
	p1, q1 := root1.Left, root1.Right
	fmt.Println(lowestCommonAncestor(root1, p1, q1).Value)
	arr2 := []int{5, 3, 8, 1, 4, 7, 9, -1, 2}
	root2 := build(arr2)
	p2, q2 := root2.Left, root2.Left.Right
	fmt.Println(lowestCommonAncestor(root2, p2, q2).Value)
}
