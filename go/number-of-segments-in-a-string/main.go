package main

import (
	"fmt"
	"strings"
)

func main() {
	s := ", , , ,        a, eaefa"
	fmt.Println(countSegments(s))
}

func countSegments(s string) (cnt int) {
	ss := strings.Trim(s, " ")
	if ss == "" {
		return 0
	}
	sss := strings.Split(ss, " ")
	for _, i := range sss {
		if len(i) != 0 && !strings.Contains(i, " ") {
			cnt += 1
		}
	}
	return cnt
}
