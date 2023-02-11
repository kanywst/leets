package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		fmt.Println(m)
		if _, ok := m[nums[i]]; ok {
			return true
		} else {
			m[nums[i]] = -1
		}
	}
	return false
}
