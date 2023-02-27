package main

import "fmt"

func main() {
	n := 0
	fmt.Println(isPowerOfTwo(n))
}

func isPowerOfTwo(n int) (ans bool) {
	if n == 1 {
		ans = true
	} else if n%2 == 1 || n == 0 {
		ans = false
	} else if n%2 == 0 {
		ans = isPowerOfTwo(n / 2)
	}
	return
}
