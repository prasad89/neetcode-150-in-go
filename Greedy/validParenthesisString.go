//go:build ignore

package main

import "fmt"

func checkValidString(s string) bool {
	min, max := 0, 0
	for _, ch := range s {
		if ch == '(' {
			min++
			max++
		} else if ch == ')' {
			min--
			max--
		} else {
			min--
			max++
		}
		if min < 0 {
			min = 0
		}
		if max < 0 {
			return false
		}
	}
	return min == 0
}

func main() {
	s1 := "((**)"
	fmt.Println(checkValidString(s1))
	s2 := "(((*)"
	fmt.Println(checkValidString(s2))
}
