package main

import "fmt"

func main() {
	nums := []int{2, 3, -1, 8, 4}
	fmt.Println(findMiddleIndex(nums))
}

func findMiddleIndex(nums []int) int {
	i, j := 0, len(nums)-1
	left, right := 0
	for {
		if left < right {
			left += nums[i]
		}
		if left > right {
			right += nums[i]
		}
	}
}
