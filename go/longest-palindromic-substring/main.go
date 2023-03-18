package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "xaabacxcabaaxcabaax"
	fmt.Println(longestPalindrome((s)))
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

func longestPalindrome(s string) string {
	tmp := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s)+1; j++ {
			ss := s[i:j]
			if IsPalindrome(ss) {
				if len(tmp) < len(ss) {
					tmp = ss
				}
			}
		}
	}
	return tmp
}
