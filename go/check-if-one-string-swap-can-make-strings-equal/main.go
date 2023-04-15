package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s1, s2 := "bank", "kanb"
	fmt.Println(areAlmostEqual(s1, s2))
}

func areAlmostEqual(s1 string, s2 string) bool {
	s1L, s2L := strings.Split(s1, ""), strings.Split(s2, "")
	sort.Strings(s1L)
	sort.Strings(s2L)
	if strings.Join(s1L, "") == strings.Join(s2L, "") {
		return true
	} else {
		return false
	}
}
