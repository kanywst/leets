package main

import (
	"fmt"
)

func main() {
	nums, k := []int{1, 2, 3, 1, 2, 3}, 2
	fmt.Println(containsNearbyDuplicate(nums, k))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if vv, ok := m[v]; ok {
			if i-vv <= k {
				return true
			}
		}
		m[v] = i
	}
	return false
}
