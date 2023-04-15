package main

import (
	"fmt"
)

func main() {
	patterns := []string{"a", "abc", "bc", "d"}
	word := "abc"
	fmt.Println(numOfStrings(patterns, word))
}

func numOfStrings(patterns []string, word string) (cnt int) {
	n := len(word)
	for _, p := range patterns {
		for i := 0; i < len(word); i++ {
			fmt.Println(p, word[i:n])
			if word[i:len(p)] == p {
				cnt += 1
			}
		}
	}
	return
}
