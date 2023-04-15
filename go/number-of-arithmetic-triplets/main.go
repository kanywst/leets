package main

import (
	"fmt"
)

func main() {
	nums, diff := []int{0, 1, 4, 6, 7, 10}, 3
	fmt.Println(arithmeticTriplets(nums, diff))
}

func arithmeticTriplets(nums []int, diff int) int {
	for i, j := range nums {
		nums[i] = j - diff
	}
	fmt.Println(nums)
	return -0
}
