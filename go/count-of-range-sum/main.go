package main

import (
	"fmt"
	"sort"
)

func main() {
	nums, lower, upper := []int{-2, 5, -1}, -2, 2
	fmt.Println(countRangeSum(nums, lower, upper))
}

func countRangeSum(nums []int, lower int, upper int) (cnt int) {
	sort.Ints(nums)
	if lower != nums[0] {
		nums = append([]int{lower}, nums...)
	}
	if upper != nums[len(nums)-1] {
		nums = append(nums, upper)
	}
	sort.Ints(nums)
	for _, i := range nums {
		if lower <= i && i <= upper {
			cnt += 1
		}
		if i > upper {
			break
		}
	}
	return
}
