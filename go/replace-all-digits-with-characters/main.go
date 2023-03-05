package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "a1c1e1"
	fmt.Println(replaceDigits(s))
}

func replaceDigits(s string) (ans string) {
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			ans += string(s[i])
		} else {
			j, _ := strconv.Atoi(string(s[i]))
			ans += string(s[i-1] + byte(j))
		}
	}
	return
}
