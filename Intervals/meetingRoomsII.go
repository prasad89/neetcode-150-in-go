//go:build ignore

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Intervals struct {
	Start int
	End   int
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[:len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return x
}

func minMeetingRooms(intervals []Intervals) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, intervals[0].End)
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start >= (*minHeap)[0] {
			heap.Pop(minHeap)
		}
		heap.Push(minHeap, intervals[i].End)
	}
	return minHeap.Len()
}

func main() {
	intervals1 := []Intervals{
		{Start: 0, End: 40},
		{Start: 5, End: 10},
		{Start: 15, End: 20},
	}
	fmt.Println(minMeetingRooms(intervals1))
	intervals2 := []Intervals{
		{Start: 4, End: 9},
	}
	fmt.Println(minMeetingRooms(intervals2))
}
