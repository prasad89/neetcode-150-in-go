//go:build ignore

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			total := nums[i] + nums[l] + nums[r]
			if total < 0 {
				l++
			} else if total > 0 {
				r--
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}
	return result
}

func main() {
	nums1 := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums1))
	nums2 := []int{}
	fmt.Println(threeSum(nums2))
}
