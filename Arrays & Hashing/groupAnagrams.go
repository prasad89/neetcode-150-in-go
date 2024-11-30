//go:build ignore

package main

import "fmt"

func groupAnagrams(words []string) [][]string {
	anagramMap := make(map[[26]int][]string)
	for _, word := range words {
		var count [26]int
		for _, ch := range word {
			count[ch-'a']++
		}
		anagramMap[count] = append(anagramMap[count], word)
	}
	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}
	return result
}

func main() {
	words1 := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	fmt.Println(groupAnagrams(words1))
	words2 := []string{"x"}
	fmt.Println(groupAnagrams(words2))
	words3 := []string{""}
	fmt.Println(groupAnagrams(words3))
}
