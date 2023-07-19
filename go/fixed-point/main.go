package main

import "fmt"

func main() {
	arr := []int{-10, -5, 0, 3, 7}
	fmt.Println(fixedPoint(arr))
}

func fixedPoint(arr []int) int {
	for i, j := range arr {
		if i == j {
			return i
		}
	}
	return -1
}
