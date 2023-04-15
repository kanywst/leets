package main

import (
	"fmt"
	"strings"
)

func main() {
	s, spaces := "LeetcodeHelpsMeLearn", []int{8, 13, 15}
	fmt.Println(addSpaces(s, spaces))
}
func addSpaces(s string, spaces []int) (ans string) {
	i := 0
	for j, k := range strings.Split(s, "") {
		if i >= len(spaces) || i >= len(s) {
			ans += s[j:]
			break
		}
		fmt.Println(j, spaces[i])
		if j == spaces[i] {
			ans += " "
			i++
		}
		ans += k
	}
	return
}
