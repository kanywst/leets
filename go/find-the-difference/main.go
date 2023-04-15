package main

import (
	"fmt"
	"strings"
)

func main() {
	s, t := "abcd", "abcde"
	fmt.Println(findTheDifference(s, t))
}

func findTheDifference(s string, t string) byte {
	for _, i := range strings.Split(s, "") {
		t = strings.Replace(t, i, "", -1)
	}
	return t[0]
}
