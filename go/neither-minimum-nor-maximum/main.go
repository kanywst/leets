package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 2, 1, 4}
	fmt.Println(findNonMinOrMax(nums))
}

func findNonMinOrMax(nums []int) int {
	if len(nums) <= 2 {
		return -1
	}
	sort.Ints(nums)
	return nums[1]
}
