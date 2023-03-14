package main

import (
	"fmt"
)

func main() {
	digits := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	l := len(digits)
	if digits[l-1] != 9 {
		digits[l-1]++
		return digits
	}
	digits[l-1] = 0
	idx := l - 2
	for ; idx >= 0; idx-- {
		if digits[idx] != 9 {
			digits[idx]++
			return digits
		} else {
			digits[idx] = 0
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}
