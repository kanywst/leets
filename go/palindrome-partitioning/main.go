package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "aab"
	fmt.Println(partition(s))
}

func IsPalindrome(s string) bool {
	l := len(s)
	for i, runeValue := range s {
		runeValue2, _ := utf8.DecodeLastRuneInString(s[:l-i])
		if runeValue != runeValue2 {
			return false
		}
	}
	return true
}

func partition(s string) [][]string {
	tmp := [][]string{}
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s)+1; j++ {
			ss := s[i:j]
			if IsPalindrome(ss) {
				tmp = append(tmp, strings.Split(ss, ""))
			}
		}
	}
	return tmp
}
