package main

import "fmt"

func main() {
	nums, k := []int{1, 2, 3, 4, 5, 6, 7}, 3
	rotate(nums, k)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	tmp := nums
	nums = append(nums[len(nums)-k:], nums[0:len(nums)-k]...)
	for i := 0; i < len(nums); i++ {
		tmp[i] = nums[i]
	}
	return
}
