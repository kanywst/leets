package main

import (
	"fmt"
	"strings"
)

func main() {
	// s := "abccbaacz"
	s := "eesll"
	fmt.Println(repeatedCharacter(s))
}

func repeatedCharacter(s string) (ans byte) {
	m := make(map[string]int)
	for _, i := range strings.Split(s, "") {
		m[i] += 1
		if m[i] >= 2 {
			ans = i[0]
			break
		}
	}
	return
}
