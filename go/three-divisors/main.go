package main

import "fmt"

func main() {
	n := 4
	fmt.Println(isThree(n))
}

func countDivisors(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}

func isThree(n int) bool {
	if countDivisors(n) == 3 {
		return true
	}
	return false
}
