package main

import "fmt"

func main() {
	s, goal := "ab", "ba"
	fmt.Println(buddyStrings(s, goal))
}

func buddyStrings(s string, goal string) (ans bool) {
	ss := ""
	for i := len(s) - 1; i >= 0; i-- {
		ss += string(s[i])
	}
	if ss == goal {
		ans = true
	}
	return
}
