package main

import "fmt"

func main() {
	nums, k := []int{34, 23, 1, 24, 75, 33, 54, 8}, 60
	fmt.Println(twoSumLessThanK(nums, k))
}

func combinations(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			combination := []int{nums[i], nums[j]}
			result = append(result, combination)
		}
	}
	return result
}

func twoSumLessThanK(nums []int, k int) int {
	diff := -1
	m := make(map[int]int)
	for _, c := range combinations(nums) {
		cc := c[0] + c[1]
		diff = k - cc
		if diff > 0 {
			m[diff] = cc
		}
	}
	for i := 0; i < k; i++ {
		if m[i] != 0 {
			return m[i]
		}
	}
	return -1
}
