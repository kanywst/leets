package main

import (
	"fmt"
	"strings"
)

func main() {
	ransomNote, magazine := "aa", "aab"
	fmt.Println(canConstruct(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) (ans bool) {
	t := ""
	for _, r := range strings.Split(ransomNote, "") {
		if strings.Contains(magazine, r) {
			t += r
		}
	}
	if ransomNote == t {
		return true
	}
	return false
}
