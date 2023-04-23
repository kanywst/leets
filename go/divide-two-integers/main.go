package main

import (
	"fmt"
	"math"
)

func main() {
	dividend, divisor := 10, 3
	fmt.Println(divide(dividend, divisor))
}

func divide(dividend int, divisor int) int {
	ans := dividend / divisor
	m := int(math.Pow(2, 31))
	if m <= ans {
		return m - 1
	}
	return dividend / divisor
}
