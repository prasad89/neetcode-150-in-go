//go:build ignore

package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func levelOrder(root *Node) []int {
	if root == nil {
		return []int{-1}
	}
	result := []int{}
	queue := []*Node{root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr != nil {
			result = append(result, curr.Value)
			queue = append(queue, curr.Left, curr.Right)
		} else {
			result = append(result, -1)
		}
	}
	for len(result) > 0 && result[len(result)-1] == -1 {
		result = result[:len(result)-1]
	}
	return result
}

func buildTree(inOrder, preOrder []int) *Node {
	if len(inOrder) == 0 || len(preOrder) == 0 {
		return nil
	}
	root := &Node{Value: preOrder[0]}
	var mid int
	for i, val := range inOrder {
		if val == preOrder[0] {
			mid = i
			break
		}
	}
	root.Left = buildTree(inOrder[:mid], preOrder[1:mid+1])
	root.Right = buildTree(inOrder[mid+1:], preOrder[mid+1:])
	if root.Right != nil && root.Right.Left == nil && root.Right.Right != nil {
		root.Right.Left = nil
	}
	return root
}

func main() {
	inOrder1 := []int{2, 1, 3, 4}
	preOrder1 := []int{1, 2, 3, 4}
	root1 := buildTree(inOrder1, preOrder1)
	fmt.Println(levelOrder(root1))
	inOrder2 := []int{1}
	preOrder2 := []int{1}
	root2 := buildTree(inOrder2, preOrder2)
	fmt.Println(levelOrder(root2))
}
