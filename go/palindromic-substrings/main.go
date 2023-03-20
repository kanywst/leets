package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "aaa"
	fmt.Println(countSubstrings(s))
}

func countSubstrings(s string) (cnt int) {
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			fmt.Println(s[i : j+1])
			if IsPalindrome(s[i : j+1]) {
				cnt += 1
			}
		}
	}
	return
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
