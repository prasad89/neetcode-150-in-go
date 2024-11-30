//go:build ignore

package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isAlpha(rune(s[l])) {
			l++
		}
		for l < r && !isAlpha(rune(s[r])) {
			r--
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlpha(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func main() {
	s1 := "Was it a car or a cat I saw?"
	fmt.Println(isPalindrome(s1))
	s2 := "tab a cat"
	fmt.Println(isPalindrome(s2))
}
