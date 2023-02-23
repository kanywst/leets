package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abacbc"
	fmt.Println(areOccurrencesEqual(s))
}

func areOccurrencesEqual(s string) bool {
	m := make(map[string]int, len(s))
	tmp := 0
	for _, ss := range strings.Split(s, "") {
		if _, ok := m[ss]; ok {
			m[ss] += 1
		} else {
			m[ss] = 0
		}
		tmp = m[ss]
	}
	for _, val := range m {
		if val != tmp {
			return false
		}
	}
	return true
}
