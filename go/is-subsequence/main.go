package main

import (
	"fmt"
	"strings"
)

func main() {
	// s, t := "abc", "ahbgdc"
	// s, t := "axc", "ahbgdc"
	// s, t := "acb", "ahbgdc"
	// s, t := "aaaaaa", "bbaaaa"
	s, t := "bb", "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	if len(t) == len(s) && s != t {
		return false
	}
	tmp := []string{}
	for _, i := range strings.Split(s, "") {
		for j, k := range strings.Split(t, "") {
			if i == k {
				tmp = append(tmp, k)
				t = t[j+1:]
				break
			}
		}
	}
	if strings.Join(tmp, "") == s {
		return true
	}
	return false
}
