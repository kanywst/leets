package main

import (
	"fmt"
	"strings"
)

func main() {
	s, k := "abcdefg", 2
	fmt.Println(reverseStr(s, k))
}

func reverseStr(s string, k int) string {
	s_head := s[0:k]
	s_tail := s[k:]
	tmp := make([]string, len(s_head))
	for i := 0; i < len(s_head); i++ {
		tmp[i] = string(s_head[len(s_head)-i-1])
	}
	return strings.Join(append(tmp, strings.Split(s_tail, "")...), "")
}
