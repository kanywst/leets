package main

import (
	"fmt"
	"sort"
)

func main() {
	nums, k := []int{1, 2, 3, 4, 5}, 3
	fmt.Println(maximizeSum(nums, k))
}

func maximizeSum(nums []int, k int) (ans int) {
	sort.Ints(nums)
	i := 0
	for i < k {
		t := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		nums = append(nums, t+1)
		ans += t
		i += 1
	}
	return
}
