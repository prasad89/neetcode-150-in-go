//go:build ignore

package main

import "fmt"

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	ptr := 0
	for {
		slow = nums[slow]
		ptr = nums[ptr]
		if slow == ptr {
			break
		}
	}
	return slow
}

func main() {
	nums1 := []int{1, 2, 3, 2, 2}
	fmt.Println(findDuplicate(nums1))
	nums2 := []int{1, 2, 3, 4, 4}
	fmt.Println(findDuplicate(nums2))
}
