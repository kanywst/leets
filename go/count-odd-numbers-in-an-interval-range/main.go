package main

import "fmt"

func main() {
	low, high := 3, 7
	fmt.Println(countOdds(low, high))
}

func countOdds(low int, high int) int {
	cnt := 0
	for i := low; i <= high; i++ {
		if i%2 == 1 {
			cnt += 1
		}
	}
	return cnt
}
