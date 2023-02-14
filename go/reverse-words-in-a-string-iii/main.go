package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}

func reverseString(s []string) string {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return strings.Join(s, "")
}

func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	ans := make([]string, len(ss))
	for i, j := range ss {
		ans[i] = reverseString(strings.Split(j, ""))
	}

	return strings.Join(ans, " ")
}
