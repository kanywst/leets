package main

import (
	"fmt"
	"strings"
)

func main() {
	word, ch := "abcdefd", 'd'
	fmt.Println(reversePrefix(word, byte(ch)))
}

func reversePrefix(word string, ch byte) string {
	ii, ws := -1, strings.Split(word, "")
	for i, w := range word {
		if ch == byte(w) {
			ii = i
			break
		}
	}
	head, tail := ws[:ii+1], ws[ii+1:]
	ans := []string{}
	for i := 0; i < len(head); i++ {
		ans = append(ans, head[len(head)-i-1])
	}
	ans = append(ans, tail...)
	return strings.Join(ans, "")
}
