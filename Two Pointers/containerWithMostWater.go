//go:build ignore

package main

import "fmt"

func maxArea(heights []int) int {
	maxWater := 0
	left, right := 0, len(heights)-1
	for left < right {
		water := (right - left) * min(heights[left], heights[right])
		maxWater = max(maxWater, water)
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return maxWater
}

func main() {
	heights1 := []int{1, 7, 2, 5, 4, 7, 3, 6}
	fmt.Println(maxArea(heights1))
	heights2 := []int{2, 2, 2}
	fmt.Println(maxArea(heights2))
}
