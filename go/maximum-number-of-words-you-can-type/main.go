package main

import (
	"fmt"
	"strings"
)

func main() {
	text, brokenLetters := "hello world", "ad"
	fmt.Println(canBeTypedWords(text, brokenLetters))
}

func canBeTypedWords(text string, brokenLetters string) (ans int) {
	m := make(map[string]bool)
	for _, b := range strings.Split(brokenLetters, "") {
		m[b] = true
	}
	for _, t := range strings.Split(text, " ") {
		for j, i := range strings.Split(t, "") {
			if _, ok := m[i]; ok {
				break
			}
			if len(t)-1 == j {
				ans += 1
			}
		}
	}
	return
}
