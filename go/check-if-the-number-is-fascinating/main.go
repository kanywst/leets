package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 783
	fmt.Println(isFascinating(n))
}

func isFascinating(n int) bool {
	s1, s2 := strconv.Itoa(n*2), strconv.Itoa(n*3)
	s := strconv.Itoa(n) + s1 + s2
	m := make(map[string]int)
	for _, ss := range strings.Split(s, "") {
		m[ss] += 1
	}
	if _, ok := m["0"]; ok {
		return false
	}
	for i := 1; i <= 9; i++ {
		if _, ok := m[strconv.Itoa(i)]; !ok {
			return false
		}
	}
	for _, v := range m {
		if v > 1 {
			return false
		}
	}
	return true
}
