package main

import "fmt"

func main() {
	num := 5
	fmt.Println(findComplement(num))
}

func findComplement(num int) int {
	x := ^0
	for x&num > 0 {
		x <<= 1
	}
	return ^x ^ num
}
