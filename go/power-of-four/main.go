package main

import (
	"fmt"
	"math"
)

func main() {
	n := 0
	fmt.Println(isPowerOfFour(n))

}

func isPowerOfFour(n int) (ans bool) {
	i := 0
	// if n == 0 {
	// 	return
	// }
	for {
		p := int(math.Pow(4, float64(i)))
		fmt.Println(p, n, p == n)
		if p == n {
			ans = true
			break
		}
		if p > n {
			ans = false
			break
		}
		i++
	}
	return
}
