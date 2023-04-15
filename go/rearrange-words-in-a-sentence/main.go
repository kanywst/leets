package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Leetcode is cool"
	fmt.Println(arrangeWords(text))
}

func arrangeWords(text string) string {
	m := make(map[string]int)
	for _, i := range strings.Split(text, " ") {
		m[i] = len(i)
	}
	n := []int{}
	for _, i := range m {
		n = append(n, i)
	}
	fmt.Println(m)
	return ""
}
