//go:build ignore

package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	i, n := 0, len(intervals)
	for i < n && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}
	fmt.Println(res)
	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	res = append(res, newInterval)
	for i < n {
		res = append(res, intervals[i])
		i++
	}
	return res
}

func main() {
	intervals1, newInterval1 := [][]int{{1, 3}, {4, 6}}, []int{2, 5}
	fmt.Println(insert(intervals1, newInterval1))
	intervals2, newInterval2 := [][]int{{1, 2}, {3, 5}, {9, 10}}, []int{6, 7}
	fmt.Println(insert(intervals2, newInterval2))
}
