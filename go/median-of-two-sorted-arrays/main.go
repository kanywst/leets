package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1, nums2 := []int{1, 2}, []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) (ans float64) {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	n := len(nums)
	if n%2 == 1 {
		ans = float64(nums[n/2])
	} else {
		ans = (float64(nums[n/2-1]) + float64(nums[n/2])) / 2.0
	}
	return
}
