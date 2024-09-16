package main

import (
	"fmt"
)

func main() {
	number, digit := "123", "3"
	fmt.Println(removeDigit(number, digit[0]))
}

func removeDigit(number string, digit byte) string {
	ans := ""
	for i, j := range number {
		if j == rune(digit) {
			ans += number[i+1:]
			break
		}
		ans += string(j)
	}
	return ans
}
