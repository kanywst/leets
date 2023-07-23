package main

import (
	"fmt"
	"strings"
)

func main() {
	words, separator := []string{"one.two.three", "four.five", "six"}, []byte(".")
	fmt.Println(splitWordsBySeparator(words, separator[0]))
}

func splitWordsBySeparator(words []string, separator byte) []string {
	s := string(separator)
	ansTmp, ans := []string{}, []string{}
	for _, w := range words {
		ansTmp = append(ansTmp, strings.Split(w, s)...)
	}
	for _, a := range ansTmp {
		if a != "" {
			ans = append(ans, a)
		}
	}
	return ans
}
