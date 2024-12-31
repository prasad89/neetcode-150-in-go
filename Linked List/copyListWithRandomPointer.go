//go:build ignore

package main

import "fmt"

type Node struct {
	Value  int
	Next   *Node
	Random *Node
}

func build(arr [][]int) *Node {
	var nodes []*Node
	for _, pair := range arr {
		nodes = append(nodes, &Node{Value: pair[0]})
	}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	for i, pair := range arr {
		if pair[1] != -1 {
			nodes[i].Random = nodes[pair[1]]
		}
	}
	return nodes[0]
}

func traverse(head *Node) {
	for current := head; current != nil; current = current.Next {
		randomValue := -1
		if current.Random != nil {
			randomValue = current.Random.Value
		}
		fmt.Printf("[%d, %d] -> ", current.Value, randomValue)
	}
	fmt.Println("nil")
}

func copyRandomList(head *Node) *Node {
	mp := map[*Node]*Node{nil: nil}
	cur := head
	for cur != nil {
		mp[cur] = &Node{Value: cur.Value}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		copy := mp[cur]
		copy.Next = mp[cur.Next]
		copy.Random = mp[cur.Random]
		cur = cur.Next
	}
	return mp[head]
}

func main() {
	arr1 := [][]int{
		{3, -1},
		{7, 3},
		{4, 0},
		{5, 1},
	}
	head1 := build(arr1)
	copiedHead1 := copyRandomList(head1)
	traverse(copiedHead1)
	arr2 := [][]int{{1, -1}, {2, 2}, {3, 2}}
	head2 := build(arr2)
	copiedHead2 := copyRandomList(head2)
	traverse(copiedHead2)
}
