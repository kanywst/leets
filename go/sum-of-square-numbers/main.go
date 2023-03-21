package main

import (
	"fmt"
	"math"
)

func main() {
	c := 5
	fmt.Println(judgeSquareSum(c))
}

func judgeSquareSum(c int) bool {
	limit := int(math.Sqrt(float64(c)))
	for i := 0; i <= limit; i++ {
		n := c - i*i
		tmp := int(math.Sqrt(float64(n)))
		if tmp*tmp == n {
			return true
		}
	}
	return false
}
