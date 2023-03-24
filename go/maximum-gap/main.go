package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{3, 6, 9, 1}
	fmt.Println(maximumGap(nums))
}

func maximumGap(nums []int) int {
	tmp := math.MinInt
	sort.Ints(nums)
	if len(nums) == 1 {
		return 0
	}
	for i := 0; i < len(nums)-1; i++ {
		n := int(math.Abs(float64(nums[i+1] - nums[i])))
		if tmp < n && n > 0 {
			tmp = n
		}
	}
	if tmp < 0 {
		return 0
	}
	return tmp
}
