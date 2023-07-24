package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 6, 2, 7, 4}
	fmt.Println(maxProductDifference(nums))
}

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]*nums[len(nums)-2] - nums[1]*nums[0]
}
