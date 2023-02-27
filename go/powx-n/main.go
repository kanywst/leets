package main

import (
	"fmt"
	"math"
)

func main() {
	x, n := 2.00000, 10
	fmt.Println(myPow(x, n))
}

func myPow(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}
