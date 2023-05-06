package main

import "fmt"

func main() {
	s, words := "iloveleetcode", []string{"i", "love", "leetcode", "apples"}
	fmt.Println(isPrefixString(s, words))
}

func isPrefixString(s string, words []string) bool {
	tmp := ""
	for _, w := range words {
		tmp += w
		if tmp == s {
			return true
		}
	}
	return false
}
