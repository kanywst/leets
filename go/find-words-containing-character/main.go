package main

import (
	"fmt"
	"strings"
)

func main() {
	words, x := []string{"leet", "code"}, byte('e')
	fmt.Println(findWordsContaining(words, x))
}

func findWordsContaining(words []string, x byte) (ans []int) {
	for i, w := range words {
		if strings.Contains(w, string(x)) {
			ans = append(ans, i)
		}
	}
	return
}
