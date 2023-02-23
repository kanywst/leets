package main

import (
	"fmt"
)

func main() {
	s := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	cnt := 0
	for i := len(s) - 1; i > -1; i-- {
		if s[i] != ' ' {
			cnt += 1
		}
		if s[i] == ' ' && cnt != 0 {
			break
		}
	}
	return cnt
}
