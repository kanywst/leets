package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

func sortedSquares(nums []int) []int {
	for i, j := range nums {
		nums[i] = j * j
	}
	sort.Ints(nums)

	return nums
}
