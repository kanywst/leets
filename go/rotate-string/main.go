package main

import "fmt"

func main() {
	s, goal := "abcde", "cdeab"
	fmt.Println(rotateString(s, goal))
}

func rotateString(s string, goal string) bool {
	for i := 0; i < len(s); i++ {
		s = s[1:] + string(s[0])
		if s == goal {
			return true
		}
	}
	return false
}
