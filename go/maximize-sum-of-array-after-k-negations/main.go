package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums, k := []int{4, 2, 3}, 1
	nums, k := []int{2, -3, -1, 5, -4}, 2
	fmt.Println(largestSumAfterKNegations(nums, k))
}

func largestSumAfterKNegations(nums []int, k int) int {
	for i := 0; i < k; i++ {
		sort.Ints(nums)
		nums[0] = -1 * nums[0]
	}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}
