//go:build ignore

package main

import "fmt"

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	count := make(map[rune]int)
	for _, ch := range s1 {
		count[ch]++
	}
	for _, ch := range s2 {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}
	return true
}

func main() {
	s1, s2 := "racecar", "carrace"
	fmt.Println(isAnagram(s1, s2))
	s3, s4 := "jar", "jam"
	fmt.Println(isAnagram(s3, s4))
}
