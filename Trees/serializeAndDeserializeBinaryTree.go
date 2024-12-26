//go:build ignore

package main

import (
	"fmt"
	"strconv"
	"strings"
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
		} else {
			curr.Left = nil
		}
		ptr++
		if ptr < len(arr) && arr[ptr] != -1 {
			curr.Right = &Node{Value: arr[ptr]}
			queue = append(queue, curr.Right)
		} else {
			curr.Right = nil
		}
		ptr++
	}
	return root
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	queue := []*Node{root}
	for len(queue) > 0 {
		level := []int{}
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			curr := queue[0]
			queue = queue[1:]
			if curr == nil {
				level = append(level, -1)
			} else {
				level = append(level, curr.Value)
				queue = append(queue, curr.Left)
				queue = append(queue, curr.Right)
			}
		}
		result = append(result, level)
	}
	return result
}

func serialize(root *Node) string {
	var serialized strings.Builder
	var preOrder func(root *Node)
	preOrder = func(root *Node) {
		if root == nil {
			serialized.WriteString("-1,")
			return
		}
		serialized.WriteString(strconv.Itoa(root.Value) + ",")
		preOrder(root.Left)
		preOrder(root.Right)
	}
	preOrder(root)
	return serialized.String()
}

func deserialize(serialized string) *Node {
	vals := strings.Split(serialized, ",")
	idx := 0
	var build func() *Node
	build = func() *Node {
		if vals[idx] == "-1" {
			idx++
			return nil
		}
		val, _ := strconv.Atoi(vals[idx])
		idx++
		root := &Node{Value: val}
		root.Left = build()
		root.Right = build()
		return root
	}
	return build()
}

func main() {
	arr1 := []int{1, 2, 3, -1, -1, 4, 5}
	root1 := build(arr1)
	serialized1 := serialize(root1)
	deserialized1 := deserialize(serialized1)
	fmt.Println(levelOrder(deserialized1))
}
