//go:build ignore

package main

import (
	"fmt"
	"strconv"
)

func encode(strs []string) string {
	var encoded string
	for _, str := range strs {
		encoded += strconv.Itoa(len(str)) + "#" + str
	}
	return encoded
}

func decode(encoded string) []string {
	var decoded []string
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(encoded[i:j])
		i = j + 1
		decoded = append(decoded, encoded[i:i+length])
		i += length
	}
	return decoded
}

func main() {
	strs1 := []string{"neet", "code", "love", "you"}
	fmt.Println(decode(encode(strs1)))
	strs2 := []string{"we", "say", ":", "yes"}
	fmt.Println(decode(encode(strs2)))
}
