//go:build ignore

package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	count := 0
	lastEnd := intervals[0][1]
	for _, interval := range intervals[1:] {
		start, end := interval[0], interval[1]
		if start < lastEnd {
			count++
			lastEnd = min(lastEnd, end)

		} else {
			lastEnd = end
		}
	}
	return count
}

func main() {
	intervals1 := [][]int{{1, 2}, {2, 4}, {1, 4}}
	fmt.Println(eraseOverlapIntervals(intervals1))
	intervals2 := [][]int{{1, 2}, {2, 4}}
	fmt.Println(eraseOverlapIntervals(intervals2))
}
