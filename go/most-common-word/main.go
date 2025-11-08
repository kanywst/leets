package main

import (
	"fmt"
	"strings"
)

func main() {
	paragraph, banned := "Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}
	fmt.Println(mostCommonWord(paragraph, banned))
}

func mostCommonWord(paragraph string, banned []string) string {
	m := make(map[string]int)
	for _, i := range strings.Split(paragraph, " ") {
		s := strings.ToLower(strings.Trim(i, "!?',;."))
		if _, ok := m[s]; ok {
			m[s] += 1
		} else {
			m[s] = 1
		}
	}

	mm := make(map[int]string)
	fmt.Println(m)
	return ""
}
