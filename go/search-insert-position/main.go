package main

import "fmt"

func main() {
	nums, target := []int{1, 3, 5, 6}, 5
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) (ans int) {
	ans = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			ans = i
			break
		}
	}
	return
}
