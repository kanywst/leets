package main

import (
	"fmt"
)

func main() {
	s := "leeetcode"
	fmt.Println(makeFancyString(s))
}

func makeFancyString(s string) string {
	ans := []byte{s[0]}
	cnt := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt += 1
			if cnt < 3 {
				ans = append(ans, s[i])
			}
		} else {
			ans = append(ans, s[i])
			cnt = 1
		}
	}
	return string(ans)
}
