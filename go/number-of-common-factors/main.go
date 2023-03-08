package main

import "fmt"

func main() {
	a, b := 12, 6
	fmt.Println(commonFactors(a, b))
}

func commonFactors(a int, b int) (cnt int) {
	n := 0
	if a < b {
		n = a
	} else {
		n = b
	}
	for i := 1; i <= n; i++ {
		fmt.Println(i)
		if a%i == 0 && b%i == 0 {
			cnt += 1
		}
	}
	return
}
