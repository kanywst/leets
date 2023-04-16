package main

import (
	"fmt"
	"strings"
)

func main() {
	order, s := "cba", "abcd"
	fmt.Println(customSortString(order, s))
}

func customSortString(order string, s string) (ans string) {
	i := 0
	for i < len(order) {
		if strings.Contains(s, string(order[i])) {
			ans += string(order[i])
			s = strings.Replace(s, string(order[i]), "", 1)
		} else {
			i += 1
		}
	}
	ans += s
	return
}
