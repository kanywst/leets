package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	// nums := []int{3, 6, 9}
	fmt.Println(minimumOperations(nums))
}

func minimumOperations(nums []int) int {
	ans := 0
	for _, i := range nums {
		temp := i % 3
		if temp == 1 {
			ans += 1
		} else if temp == 2 {
			ans += 1
		}
	}
	return ans
}
