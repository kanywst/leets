package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "AbCdEfGhIjK"
	fmt.Println(greatestLetter(s))
}

func greatestLetter(s string) string {
	m1, m2 := make(map[rune]bool), make(map[rune]bool)
	for _, i := range s {
		if i >= 65 && i <= 90 {
			m1[i] = true
		}
		if i >= 97 && i <= 122 {
			m2[i] = true
		}
	}
	ans := []string{}
	for k, _ := range m1 {
		if _, ok := m2[k+32]; ok {
			ans = append(ans, string(k))
		}
	}
	sort.Strings(ans)
	if len(ans) == 0 {
		return ""
	}
	return ans[len(ans)-1]
}
