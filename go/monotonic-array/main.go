package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 2, 3}
	fmt.Println(isMonotonic(nums))
}

func isMonotonic(nums []int) (ans bool) {
	f, ans := false, true
	if nums[0] < nums[len(nums)-1] {
		f = true
	}
	for i := 0; i < len(nums)-1; i++ {
		if f && nums[i] > nums[i+1] {
			ans = false
			break
		} else if !f && nums[i] < nums[i+1] {
			ans = false
			break
		}
	}
	return
}
