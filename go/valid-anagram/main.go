package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s, t := "anagram", "nagaram"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ss := strings.Split(s, "")
	tt := strings.Split(t, "")
	sort.Slice(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})
	sort.Slice(tt, func(i, j int) bool {
		return tt[i] < tt[j]
	})
	for i := 0; i < len(ss); i++ {
		if ss[i] != tt[i] {
			return false
		}
	}
	return true
}
