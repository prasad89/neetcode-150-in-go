//go:build ignore

package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	frequency := make([][]int, len(nums)+1)
	for num, cnt := range count {
		frequency[cnt] = append(frequency[cnt], num)
	}
	result := make([]int, 0, k)
	for i := len(frequency) - 1; i >= 0 && len(result) < k; i-- {
		for _, num := range frequency[i] {
			result = append(result, num)
			if len(result) == k {
				return result
			}
		}
	}
	return result
}

func main() {
	nums1, k1 := []int{1, 2, 2, 3, 3, 3}, 2
	fmt.Println(topKFrequent(nums1, k1))
	nums2, k2 := []int{7, 7}, 1
	fmt.Println(topKFrequent(nums2, k2))
}
