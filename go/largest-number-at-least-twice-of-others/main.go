package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 6, 1, 0}
	fmt.Println(dominantIndex(nums))
}

func dominantIndex(nums []int) int {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(nums)
	m := nums[len(nums)-1]
	nums = nums[:len(nums)-1]
	for _, i := range nums {
		if i*2 > m {
			return -1
		}
	}
	ans := -1
	for i, j := range tmp {
		if j == m {
			ans = i
		}
	}
	return ans
}
