package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(countDigits(123))
}
func countDigits(num int) int {
	s := strconv.Itoa(num)
	c := 0
	for _, y := range strings.Split(s, "") {
		i, _ := strconv.Atoi(y)
		if num%i == 0 {
			c += 1
		}
	}
	return c
}
