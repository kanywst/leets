package main

import (
	"fmt"
	"strings"
)

func main() {
	words, left, right := []string{"are", "amy", "u"}, 0, 2
	fmt.Println(vowelStrings(words, left, right))

}

func vowelStrings(words []string, left int, right int) (cnt int) {
	ans := "aeiou"
	for _, w := range words[left : right+1] {
		if strings.Contains(ans, string(w[0])) && strings.Contains(ans, string(w[len(w)-1])) {
			cnt += 1
		}
	}
	return
}
