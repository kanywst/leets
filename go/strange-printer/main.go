package main

import (
	"fmt"
)

func main() {
	// s := "aaabbb"
	s := "aba"
	fmt.Println(strangePrinter(s))
}

func strangePrinter(s string) int {
	tmp := s[0]
	ss := ""
	for i := 1; i < len(s); i++ {
		if tmp != s[i] {
			fmt.Println(s[:i], s[i:])
			ss += string(s[i-1])
			tmp = s[i]
		}
	}
	fmt.Println(ss)
	return len(ss)
}
