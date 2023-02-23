package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num := 38
	fmt.Println(addDigits(num))
}

func addDigits(num int) int {
	for len(strconv.Itoa(num)) > 1 {
		ans := 0
		for _, i := range strings.Split(strconv.Itoa(num), "") {
			ii, _ := strconv.Atoi(i)
			ans += int(ii)
		}
		num = ans
	}
	return num
}
