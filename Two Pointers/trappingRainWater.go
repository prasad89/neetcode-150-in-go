//go:build ignore

package main

import "fmt"

func trap(height []int) int {
	totalWater := 0
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	for left < right {
		if leftMax < rightMax {
			left++
			leftMax = max(leftMax, height[left])
			totalWater += leftMax - height[left]
		} else {
			right--
			rightMax = max(rightMax, height[right])
			totalWater += rightMax - height[right]
		}
	}
	return totalWater
}

func main() {
	height := []int{0, 2, 0, 3, 1, 0, 1, 3, 2, 1}
	fmt.Println(trap(height))
}
