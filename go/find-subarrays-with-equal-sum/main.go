package main

import "fmt"

func main() {
	nums := []int{4, 2, 4}
	fmt.Println(findSubarrays(nums))
}

func findSubarrays(nums []int) bool {
	m := make(map[int]bool)
	for i := 0; i < len(nums)-1; i++ {
		sum := nums[i] + nums[i+1]
		if _, ok := m[sum]; ok {
			return true
		}
		m[sum] = true
	}
	return false
}
