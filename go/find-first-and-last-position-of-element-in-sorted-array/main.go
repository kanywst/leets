package main

import "fmt"

func main() {
	nums, target := []int{5, 7, 7, 8, 8, 10}, 8
	// nums, target := []int{5, 7, 7, 8, 8, 10}, 6
	fmt.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	ns := []int{}
	for i, j := range nums {
		if j == target {
			ns = append(ns, i)
		}
	}
	if len(ns) >= 2 {
		return []int{ns[0], ns[len(ns)-1]}
	}
	if len(ns) >= 1 {
		return []int{ns[0], ns[0]}
	}
	return []int{-1, -1}
}
