package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 2
	fmt.Println(isHappy(n))
}

func isHappy(n int) (ans bool) {
	cal := 0
	for {
		tmp := strings.Split(strconv.Itoa(n), "")
		if len(tmp) < 2 {
			ans = false
			break
		} else {
			for _, i := range tmp {
				a, _ := strconv.Atoi(i)
				cal += a * a
				if cal == 1 {
					ans = true
				}
			}
			n, cal = cal, 0
		}
	}
	return
}
