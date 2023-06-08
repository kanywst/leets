package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{10, 20, 30, 5, 10, 50}
	fmt.Println(maxAscendingSum(nums))
}

func ascendingCheck(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}
	return true
}

func sum(x []int) (ans int) {
	for _, i := range x {
		ans += i
	}
	return
}

func maxAscendingSum(nums []int) (ans int) {
	tmp := []int{}
	start := 0
	for i := 1; i < len(nums); i++ {
		if !ascendingCheck(nums[start : i+1]) {
			tmp = append(tmp, sum(nums[start:i]))
			start = i
		}
	}
	tmp = append(tmp, sum(nums[start:]))
	sort.Ints(tmp)
	return tmp[len(tmp)-1]
}
