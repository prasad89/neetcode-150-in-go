//go:build ignore

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1

	for l < r {
		if nums[l]+nums[r] < target {
			l++
		} else if nums[l]+nums[r] > target {
			r--
		} else {
			return []int{l + 1, r + 1}
		}
	}

	return []int{}
}

func main() {
	nums := []int{1, 2, 3, 4}
	target := 3
	fmt.Println(twoSum(nums, target))
}
