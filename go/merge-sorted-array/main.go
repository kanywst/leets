package main

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	tmp := append(nums1[:m], nums2[:n]...)
	sort.Ints(tmp)
	nums1 = []int{}
	nums1 = append(nums1, tmp...)
}
