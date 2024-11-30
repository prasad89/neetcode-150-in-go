//go:build ignore

package main

import "fmt"

func hasDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = true
	}
	return false
}

func main() {
	nums1 := []int{1, 2, 3, 3}
	fmt.Println(hasDuplicate(nums1))
	nums2 := []int{1, 2, 3, 4}
	fmt.Println(hasDuplicate(nums2))
}
