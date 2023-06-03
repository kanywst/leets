package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 6, 9, 1}
	fmt.Println(maximumGap(nums))
}

func maximumGap(nums []int) (ans int) {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		t := nums[i+1] - nums[i]
		if ans < t {
			ans = t
		}
	}
	return
}
