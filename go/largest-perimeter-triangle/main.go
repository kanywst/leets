package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{2, 1, 2}
	nums := []int{1, 2, 1, 10}
	fmt.Println(largestPerimeter(nums))
}

func largestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	if nums[0]+nums[1] > nums[2] {
		return nums[0] + nums[1] + nums[2]
	} else {
		return 0
	}
}
