package main

import "fmt"

func main() {
	nums, k := []int{3, 1, 2, 2, 2, 1, 3}, 2
	fmt.Println(countPairs(nums, k))
}

func countPairs(nums []int, k int) (cnt int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && (i*j)%k == 0 {
				cnt += 1
			}
		}
	}
	return
}
