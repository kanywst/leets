package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 10, 6, 7, -7, 1}
	fmt.Println(findMaxK(nums))
}

func findMaxK(nums []int) int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n] += 1
	}
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		if _, ok := m[-1*nums[i]]; ok {
			return nums[i]
		}
	}
	return -1
}
