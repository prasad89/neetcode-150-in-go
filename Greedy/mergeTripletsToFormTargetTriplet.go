//go:build ignore

package main

import "fmt"

func mergeTriplets(triplets [][]int, target []int) bool {
	set := map[int]struct{}{}
	for _, triplet := range triplets {
		if target[0] < triplet[0] || target[1] < triplet[1] || target[2] < triplet[2] {
			continue
		}
		for idx, num := range triplet {
			if target[idx] == num {
				set[idx] = struct{}{}
			}
		}
	}
	return len(set) == 3
}

func main() {
	triplets1 := [][]int{{1, 2, 3}, {7, 1, 1}}
	target1 := []int{7, 2, 3}
	fmt.Println(mergeTriplets(triplets1, target1))
	triplets2 := [][]int{{2, 5, 6}, {1, 4, 4}, {5, 7, 5}}
	target2 := []int{5, 4, 6}
	fmt.Println(mergeTriplets(triplets2, target2))
}
