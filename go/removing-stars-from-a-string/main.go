package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "leet**cod*e"
	fmt.Println(removeStars(s))
}

func removeStars(s string) string {
	ans := []string{}
	for _, j := range strings.Split(s, "") {
		if j == "*" {
			ans = ans[:len(ans)-1]
		} else {
			ans = append(ans, j)
		}
	}
	return strings.Join(ans, "")
}
