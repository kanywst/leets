package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}

func subsets(nums []int) (ans [][]int) {
	ans = append(ans, []int{})
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			fmt.Println(nums[i], nums[j])
		}
	}
	return
}
