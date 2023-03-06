package main

import (
	"fmt"
	"math"
)

func main() {
	num := 14
	fmt.Println(isPerfectSquare(num))
}

func isPerfectSquare(num int) (ans bool) {
	n := math.Sqrt(float64(num))
	if int(n)*int(n) == num {
		ans = true
	} else {
		ans = false
	}
	return
}
