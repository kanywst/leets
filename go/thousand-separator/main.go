package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 1234
	fmt.Println(thousandSeparator(n))
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func thousandSeparator(n int) string {
	s := strconv.Itoa(n)
	if len(s) <= 3 {
		return s
	}
	ss := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		ss = append(ss, string(s[i]))
		if (len(s)-i)%3 == 0 {
			ss = append(ss, ".")
		}
	}
	tmp := reverse(strings.Join(ss, ""))
	if string(tmp[0]) == "." {
		return tmp[1:]
	}
	return tmp
}
