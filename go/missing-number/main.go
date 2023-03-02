package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 1}
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) (ans int) {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		fmt.Println(i+nums[0], nums[i])
		if i != nums[i] {
			ans = i
			return
		}
	}
	ans = nums[len(nums)-1] + 1
	return
}
