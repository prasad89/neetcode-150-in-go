//go:build ignore

package main

import "fmt"

func partitionLabels(s string) []int {
	frequency := make([]int, 26)
	for i, ch := range s {
		frequency[ch-'a'] = i
	}
	partition := []int{}
	start, end := 0, 0
	for i, ch := range s {
		end = max(end, frequency[ch-'a'])
		if i == end {
			partition = append(partition, end-start+1)
			start = end + 1
		}
	}
	return partition
}

func main() {
	s1 := "xyxxyzbzbbisl"
	fmt.Println(partitionLabels(s1))
	s2 := "abcabc"
	fmt.Println(partitionLabels(s2))
}
