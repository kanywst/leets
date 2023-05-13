package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 5, 0, 3, 5}
	fmt.Println(minimumOperations(nums))
}

func zeroCheck(nums []int) bool {
	for _, n := range nums {
		if n != 0 {
			return false
		}
	}
	return true
}

func minimumOperations(nums []int) (cnt int) {
	for x := 1; x <= 100; x++ {
		sort.Ints(nums)
		for _, i := range nums {
			if i != 0 {
				x = i
				break
			}
		}
		if zeroCheck(nums) {
			break
		}
		for i, _ := range nums {
			nums[i] -= x
			if nums[i] <= 0 {
				nums[i] = 0
			}
		}
		cnt += 1
	}
	return
}
