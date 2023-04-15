package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "leet**cod*e"
	fmt.Println(removeStars(s))
}

func removeStars(s string) string {
	for {
		ind := strings.Index(s, "*")
		if ind == -1 {
			break
		}
		s = s[:ind-1] + s[ind+1:]
	}
	t := s
	return t
}
