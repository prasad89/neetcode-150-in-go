//go:build ignore

package main

import "fmt"

func jump(nums []int) int {
	jumps, farthest, current := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == current {
			current = farthest
			jumps++
		}
	}
	return jumps
}

func main() {
	nums1 := []int{2, 4, 1, 1, 1, 1}
	fmt.Println(jump(nums1))
	nums2 := []int{2, 1, 2, 1, 0}
	fmt.Println(jump(nums2))
}
