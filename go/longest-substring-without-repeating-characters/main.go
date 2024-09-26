package main

import "fmt"

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	ans := []rune{}
	for i, j := range s {
		if _, ok := m[j]; !ok {
			m[j] = i
			ans = append(ans, j)
		}
	}
	return len(ans)
}
