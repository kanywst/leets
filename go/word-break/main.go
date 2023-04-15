package main

import "fmt"

func main() {
	s, wordDict := "applepenapple", []string{"apple", "pen"}
	fmt.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]int)
	for _, w := range wordDict {
		fmt.Println(s)
		for i := 0; i < len(s)-len(w)+1; i++ {
			if w == s[i:len(w)] {
				m[w] += 1
				s = s[:i] + s[len(w):]
			}
		}
	}
	fmt.Println(m)
	if len(m) == len(wordDict) {
		return true
	}
	return false
}
