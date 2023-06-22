package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aaabc"
	fmt.Println(minimizedStringLength(s))
}

func uniqs(arr []string) (uniq []string) {
	m := make(map[string]bool)

	for _, ele := range arr {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}
	return
}

func minimizedStringLength(s string) int {
	return len(uniqs(strings.Split(s, "")))
}
