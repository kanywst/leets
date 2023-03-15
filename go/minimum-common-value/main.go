package main

import (
	"fmt"
)

func main() {
	nums1, nums2 := []int{1, 2, 3, 6}, []int{2, 3, 4, 5}
	fmt.Println(getCommon(nums1, nums2))
}

func getCommon(nums1 []int, nums2 []int) int {
	for _, i := range nums1 {
		for _, j := range nums2 {
			if i == j {
				return i
			}
		}
	}
	return -1
}
