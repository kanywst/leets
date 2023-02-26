package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	x := 1534236469
	fmt.Println(reverse(x))
}

func reverse(x int) (a int) {
	s := strconv.Itoa(x)
	f := false
	ans := []string{}
	if s[0] == '-' {
		f = true
	}
	for i, _ := range s {
		ans = append(ans, string(s[len(s)-i-1]))
	}
	aa := strings.Join(ans, "")
	if f {
		aa = aa[:len(s)-1]
	}
	a, _ = strconv.Atoi(aa)
	if f {
		a *= -1
	}
	if -1*int(math.Pow(2, 31)) > a || a > int(math.Pow(2, 31))-1 {
		return 0
	}
	return
}
