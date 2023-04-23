package main

import "fmt"

func main() {
	n := 7
	fmt.Println(sumOfMultiples(n))
}

func sumOfMultiples(n int) (ans int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			ans += i
		}
	}
	return
}
