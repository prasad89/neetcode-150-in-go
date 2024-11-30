//go:build ignore

package main

import "fmt"

func longestConsecutive(nums []int) int {
	set := map[int]struct{}{}
	for _, num := range nums {
		set[num] = struct{}{}
	}
	longest := 0
	for num := range set {
		if _, found := set[num-1]; !found {
			length := 1
			for {
				if _, found := set[num+length]; found {
					length++
				} else {
					break
				}
			}
			longest = max(length, longest)
		}
	}
	return longest
}

func main() {
	nums1 := []int{2, 20, 4, 10, 3, 4, 5}
	fmt.Println(longestConsecutive(nums1))
	nums2 := []int{0, 3, 2, 5, 4, 6, 1, 1}
	fmt.Println(longestConsecutive(nums2))
}
