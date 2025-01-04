//go:build ignore

package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	count := make(map[int]int)
	for _, card := range hand {
		count[card]++
	}
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for key := range count {
		heap.Push(minHeap, key)
	}
	for minHeap.Len() > 0 {
		firstCard := (*minHeap)[0]
		for i := firstCard; i < firstCard+groupSize; i++ {
			if count[i] == 0 {
				return false
			}
			count[i]--
			if count[i] == 0 {
				if i != (*minHeap)[0] {
					return false
				}
				heap.Pop(minHeap)
			}
		}
	}
	return true
}

func main() {
	hand1, groupSize1 := []int{1, 2, 4, 2, 3, 5, 3, 4}, 4
	fmt.Println(isNStraightHand(hand1, groupSize1))
	hand2, groupSize2 := []int{1, 2, 3, 3, 4, 5, 6, 7}, 4
	fmt.Println(isNStraightHand(hand2, groupSize2))
}
