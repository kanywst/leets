package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1, nums2 := []int{2, 6, 4}, []int{9, 7, 5}
	fmt.Println(addedInteger(nums1, nums2))
}

func asc_sort(list []int) []int {
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	return list
}
func addedInteger(nums1 []int, nums2 []int) int {
	asc_nums1 := asc_sort(nums1)
	asc_nums2 := asc_sort(nums2)
	return asc_nums2[0] - asc_nums1[0]
}
