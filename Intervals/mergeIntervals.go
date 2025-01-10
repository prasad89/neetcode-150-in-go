//go:build ignore

package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	res = append(res, intervals[0])
	for _, interval := range intervals {
		start, end := interval[0], interval[1]
		lastEnd := res[len(res)-1][1]
		if start <= lastEnd {
			res[len(res)-1][1] = max(lastEnd, end)
		} else {
			res = append(res, interval)
		}
	}
	return res
}

func main() {
	intervals1 := [][]int{{1, 3}, {1, 5}, {6, 7}}
	fmt.Println(merge(intervals1))
	intervals2 := [][]int{{1, 2}, {2, 3}}
	fmt.Println(merge(intervals2))
}
