package main

import (
	"fmt"
)

func main() {
	num := "50"
	fmt.Println(removeTrailingZeros(num))
}

func removeTrailingZeros(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		s := num[:i+1]
		if s[len(s)-1] != '0' {
			return s
		}
	}
	return num
}
