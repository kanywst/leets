package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "l|*e*et|c**o|*de|"
	fmt.Println(countAsterisks(s))
}

func countAsterisks(s string) int {
	cnt_max := 0
	for _, ss := range strings.Split(s, "|") {
		cnt := 0
		for _, sss := range ss {
			if string(sss) == "*" {
				cnt += 1
			}
		}
		if cnt_max < cnt {
			cnt_max = cnt
		}
	}
	return cnt_max
}
