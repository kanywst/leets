package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "cbacdcbc"
	fmt.Println(removeDuplicateLetters(s))
}

func removeDuplicateLetters(s string) string {
	m, tmp := make(map[string]int), []string{}
	for _, i := range strings.Split(s, "") {
		if _, ok := m[i]; ok {
			m[i] += 1
		} else {
			m[i] = 0
		}
	}
	for k, _ := range m {
		tmp = append(tmp, k)
	}
	return strings.Join(tmp, "")
}
