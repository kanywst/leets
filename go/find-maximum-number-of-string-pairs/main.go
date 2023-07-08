package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"cd", "ac", "dc", "ca", "zz"}
	fmt.Println(maximumNumberOfStringPairs(words))
}

func reverseString(s []string) string {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return strings.Join(s, "")
}

func maximumNumberOfStringPairs(words []string) int {
	count := make(map[string]int)
	pairs := 0

	for _, word := range words {
		reversed := reverseString(strings.Split(word, ""))
		pairs += count[reversed]
		count[word]++
	}

	return pairs
}
