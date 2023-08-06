package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "string"
	fmt.Println(finalString(s))
}

func reverseString(ss string) string {
	s := strings.Split(ss, "")
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return strings.Join(s, "")
}

func finalString(s string) (ans string) {
	for _, i := range strings.Split(s, "") {
		if i == "i" {
			ans = reverseString(ans)
		} else {
			ans += i
		}
	}
	return
}
