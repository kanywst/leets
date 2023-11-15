package main

import "fmt"

func main() {
	nums, target := []int{-1, 1, 2, 3, 1}, 2
	fmt.Println(countPairs(nums, target))
}

func countPairs(nums []int, target int) (cnt int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i < j && nums[i]+nums[j] < target {
				cnt += 1
			}
		}
	}
	return
}
