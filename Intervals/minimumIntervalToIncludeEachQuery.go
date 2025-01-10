//go:build ignore

package main

import "fmt"

func minInterval(intervals [][]int, queries []int) []int {
	res := make([]int, len(queries))
	for i, query := range queries {
		curr := -1
		for _, interval := range intervals {
			start, end := interval[0], interval[1]
			if start <= query && query <= end {
				if curr == -1 || end-start+1 < curr {
					curr = end - start + 1
				}
			}
		}
		res[i] = curr
	}
	return res
}

func main() {
	intervals1 := [][]int{{1, 3}, {2, 3}, {3, 7}, {6, 6}}
	queries1 := []int{2, 3, 1, 7, 6, 8}
	fmt.Println(minInterval(intervals1, queries1))
}
