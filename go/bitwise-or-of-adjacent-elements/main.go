package main

import "fmt"

func main() {
	nums := []int{1, 3, 7, 15}
	fmt.Println(orArray(nums))
}

func orArray(nums []int) []int {
	result := make([]int, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		result[i] = nums[i+1] | nums[i]
	}
	return result
}
