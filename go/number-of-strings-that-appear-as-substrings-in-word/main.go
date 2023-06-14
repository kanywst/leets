package main

import (
	"fmt"
	"strings"
)

func main() {
	patterns, word := []string{"a", "abc", "bc", "d"}, "abc"
	fmt.Println(numOfStrings(patterns, word))
}

func numOfStrings(patterns []string, word string) (cnt int) {
	for _, p := range patterns {
		if strings.Contains(word, p) {
			cnt += 1
		}
	}
	return
}
