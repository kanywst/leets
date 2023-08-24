package main

import (
	"fmt"
)

func main() {
	// words, s := []string{"alice", "bob", "charlie"}, "abc"
	words, s := []string{"an", "apple"}, "a"
	fmt.Println(isAcronym(words, s))
}

func isAcronym(words []string, s string) bool {
	if len(words) > len(s) {
		return false
	}
	for i, j := range s {
		if i >= len(words) || words[i][0] != byte(j) {
			return false
		}
	}
	return true
}
