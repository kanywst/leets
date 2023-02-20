package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	ans := make([]string, len(s))
	rev := make([]string, len(s))
	s = strings.ToLower(s)
	for i := range s {
		tmp1 := int(s[i])
		tmp2 := int(s[len(s)-i-1])

		if (tmp1 >= 48 && tmp1 <= 57) || (tmp1 >= 97 && tmp1 <= 122) {
			ans[i] = string(tmp1)
		}
		if (tmp2 >= 48 && tmp2 <= 57) || (tmp2 >= 97 && tmp2 <= 122) {
			rev[i] = string(tmp2)
		}
	}
	if strings.Join(ans, "") == strings.Join(rev, "") {
		return true
	}
	return false
}
