//go:build ignore

package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	maxSum, currSum := math.MinInt, 0
	start, end, tmpStart := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if currSum+nums[i] < nums[i] {
			currSum = nums[i]
			tmpStart = i
		} else {
			currSum += nums[i]
		}
		if maxSum < currSum {
			maxSum = currSum
			start = tmpStart
			end = i + 1
		}
	}
	fmt.Println(nums[start:end])
	return maxSum
}

func main() {
	nums1 := []int{2, -3, 4, -2, 2, 1, -1, 4}
	fmt.Println(maxSubArray(nums1))
	nums2 := []int{-1}
	fmt.Println(maxSubArray(nums2))
}
