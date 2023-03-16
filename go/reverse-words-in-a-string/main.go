package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "a good   example"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	ss, tmp := strings.Split(strings.Trim(s, " "), " "), []string{}
	for i := len(ss) - 1; i >= 0; i-- {
		if ss[i] != " " {
			tmp = append(tmp, ss[i])
		}
	}
	return strings.Join(strings.Fields(strings.Join(tmp, " ")), " ")
}
