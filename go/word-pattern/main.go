package main

import (
	"fmt"
	"strings"
)

func main() {
	pattern, s := "abba", "dog dog dog dog"
	fmt.Println(wordPattern(pattern, s))
}

func wordPattern(pattern string, s string) bool {
	if len(pattern) != len(strings.Split(s, " ")) {
		return false
	}
	m := make(map[string]string)
	tmp := ""
	for i, ss := range strings.Split(s, " ") {
		if _, ok := m[string(pattern[i])]; ok {
			if m[string(pattern[i])] != ss {
				return false
			}
		}
		m[string(pattern[i])] = ss
	}
	for _, v := range m {
		if strings.Contains(tmp, v) {
			return false
		}
		tmp += v
	}
	return true
}
