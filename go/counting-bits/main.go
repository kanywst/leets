package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 2
	fmt.Println(countBits(n))
}

func countBits(n int) []int {
	for i := 0; i < n+1; i++ {
		a, _ := strconv.ParseInt(strconv.Itoa(i), 2, 255)
		fmt.Println(a)
	}
	return []int{}
}
