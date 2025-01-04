//go:build ignore

package main

import "fmt"

func canJump(nums []int) bool {
	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		if i > maxIndex {
			return false
		}
		maxIndex = max(maxIndex, i+nums[i])
		if maxIndex >= len(nums)-1 {
			return true
		}
	}
	return true
}

func main() {
	nums1 := []int{1, 2, 0, 1, 0}
	fmt.Println(canJump(nums1))
	nums2 := []int{1, 2, 1, 0, 1}
	fmt.Println(canJump(nums2))
}
