package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := "35427"
	fmt.Println(largestOddNumber(num))
}

func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		t := num[:i+1]
		tt, _ := strconv.Atoi(string(t[len(t)-1]))
		if tt%2 != 0 {
			return num[:i+1]
		}
	}
	return ""
}
