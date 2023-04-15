package main

import (
	"fmt"
)

func main() {
	s, t := "abc", "ahbgdc"
	// s, t := "axc", "ahbgdc"
	// s, t := "", "ahbgdc"
	// s, t := "acb", "ahbgdc"
	// s, t := "aaaaaa", "bbaaaa"
	fmt.Println(isSubsequence(s, t))
}

func isSubsequence(s string, t string) (ans bool) {
	if s == "" {
		return true
	}
	cnt := 0
	for i := 0; i < len(s); i++ {
		tmp := s[i]
		for j := i; j < len(t); j++ {
			fmt.Println(string(tmp), string(t[j]), i, j, ans)
			if tmp == t[j] {
				ans = true
				i = j
				cnt += 1
				break
			} else {
				ans = false
			}
		}
		fmt.Println("cnt:", cnt)
		if !ans {
			break
		}
	}
	fmt.Println("ans:", ans)
	if cnt > len(s) {
		ans = false
	}
	return
}
