package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "race car"
	fmt.Println(reverseVowels(s))
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseVowels(s string) string {
	ind := []int{}
	tmp := []string{}
	cnt := 0
	for j, i := range s {
		if i == 'a' || i == 'i' || i == 'u' || i == 'e' || i == 'o' || i == 'A' || i == 'I' || i == 'U' || i == 'E' || i == 'O' {
			cnt += 1
			ind = append(ind, j)
			tmp = append(tmp, string(i))
		}
	}
	if cnt < 2 {
		return s
	}
	ss := reverse(strings.Join(tmp, ""))
	ans := []string{}
	for i := 0; i < len(s); i++ {
		if len(ind) > 0 {
			if ind[0] == i {
				ans = append(ans, string(ss[0]))
				ss = ss[1:]
				ind = ind[1:]
			} else {
				ans = append(ans, string(s[i]))
			}
		} else {
			ans = append(ans, string(s[i]))
		}
	}
	return strings.Join(ans, "")
}
