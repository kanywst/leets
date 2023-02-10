package main

import "fmt"

func main() {
	n := 6
	fmt.Println(smallestEvenMultiple(n))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func smallestEvenMultiple(n int) (cal int) {
	cal = n * 2 / gcd(n, 2)
	return
}
