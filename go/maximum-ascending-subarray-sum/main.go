package main

import "fmt"

func main() {
	// nums := []int{10, 20, 30, 5, 10, 50}
	// nums := []int{1}
	nums := []int{10, 20, 30, 40, 50}
	fmt.Println(maxAscendingSum(nums))
}

func sum(x []int) (ans int) {
	for _, i := range x {
		ans += i
	}
	return
}

func maxAscendingSum(nums []int) (ans int) {
	if len(nums) < 2 {
		ans = sum(nums)
		return
	}
	tmp, ind := nums[len(nums)-1], 0
	for i := len(nums) - 2; i >= 0; i-- {
		if tmp < nums[i] {
			ind = i + 1
			break
		} else {
			tmp = nums[i]
		}
	}
	ans = sum(nums[ind:])
	return
}
