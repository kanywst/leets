package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	x := -121
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) (ans bool) {
	tmp, ans := []string{}, false
	s := strings.Split(strconv.Itoa(x), "")
	for i := 0; i < len(s); i++ {
		tmp = append(tmp, s[len(s)-i-1])
	}
	n, _ := strconv.Atoi(strings.Join(tmp, ""))
	if x == n {
		ans = true
	}
	return
}
