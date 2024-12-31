//go:build ignore

package main

import "fmt"

type Node struct {
	Key, Value int
	Prev, Next *Node
}

type LRUCache struct {
	Capacity    int
	Cache       map[int]*Node
	Left, Right *Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		Capacity: capacity,
		Cache:    make(map[int]*Node),
		Left:     &Node{},
		Right:    &Node{},
	}
	lru.Left.Next = lru.Right
	lru.Right.Prev = lru.Left
	return lru
}

func (this *LRUCache) remove(node *Node) {
	next, prev := node.Next, node.Prev
	prev.Next = next
	next.Prev = prev
}

func (this *LRUCache) insert(node *Node) {
	prev, next := this.Right.Prev, this.Right
	prev.Next = node
	next.Prev = node
	node.Next = next
	node.Prev = prev
}

func (this *LRUCache) get(key int) int {
	if node, exist := this.Cache[key]; exist {
		this.remove(node)
		this.insert(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) put(key, val int) {
	if node, exist := this.Cache[key]; exist {
		this.remove(node)
		delete(this.Cache, key)
	}
	node := &Node{Key: key, Value: val}
	this.Cache[key] = node
	this.insert(node)
	if len(this.Cache) > this.Capacity {
		lru := this.Left.Next
		this.remove(lru)
		delete(this.Cache, lru.Key)
	}
}

func main() {
	commands := []string{"LRUCache", "put", "get", "put", "put", "get", "get"}
	args := [][]int{{2}, {1, 10}, {1}, {2, 20}, {3, 30}, {2}, {1}}
	var lru LRUCache
	result := []interface{}{nil}
	for i := 0; i < len(commands); i++ {
		switch commands[i] {
		case "LRUCache":
			lru = Constructor(args[i][0])
			result = append(result, nil)
		case "put":
			lru.put(args[i][0], args[i][1])
			result = append(result, nil)
		case "get":
			result = append(result, lru.get(args[i][0]))
		}
	}
	fmt.Println(result[1:])
}
