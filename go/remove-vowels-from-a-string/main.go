package main

import (
	"fmt"
)

func main() {
	s := "leetcodeisacommunityforcoders"
	fmt.Println(removeVowels(s))
}

func removeVowels(s string) string {
	ans := ""
	for _, i := range s {
		if i != 'a' && i != 'e' && i != 'i' && i != 'o' && i != 'u' {
			ans += string(i)
		}
	}
	return ans
}
