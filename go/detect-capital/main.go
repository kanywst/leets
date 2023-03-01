package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "FlaG"
	fmt.Println(detectCapitalUse(word))
}

func detectCapitalUse(word string) (ans bool) {
	usa, leetcode := strings.ToUpper(word), strings.ToLower(word)
	google := strings.Title(leetcode)
	if usa == word || leetcode == word || google == word {
		ans = true
	}
	return
}
