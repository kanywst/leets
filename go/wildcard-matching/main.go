package main

import (
	"fmt"
	"regexp"
)

func main() {
	s, p := "aa", "a"
	// s, p := "aa", "a*"
	fmt.Println(isMatch(s, p))
}

func isMatch(s string, p string) bool {
	fmt.Println(regexp.MustCompile(p).MatchString(s))
	return false
}
