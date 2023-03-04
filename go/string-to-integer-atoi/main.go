package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// s := "    4193 with words"
	// s := "-91283472332"
	// s := "3.14159"
	// s := "  -0012a42"
	s := "-5-"
	fmt.Println(myAtoi(s))
}

func myAtoi(s string) (ans int) {
	tmp := strings.Split(strings.Trim(s, " "), " ")
	ss := strings.Split(tmp[0], ".")
	sss := ""
	for _, j := range ss[0] {
		if (j == '+' || j == '-') && len(sss) > 0 {
			break
		}
		if (j != '+') && (j != '-') && (j < 48 || j > 57) {
			break
		} else {
			fmt.Println(string(j))
			sss += string(j)
		}
	}
	ans, _ = strconv.Atoi(sss)
	if ans < math.MinInt32 {
		ans = math.MinInt32
	}
	if ans > math.MaxInt32 {
		ans = math.MaxInt32
	}
	return
}
