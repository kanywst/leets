package main

import "fmt"

func main() {
	nums1, nums2 := []int{1, 2, 3, 3}, []int{1, 1, 2, 2}
	fmt.Println(findDifference(nums1, nums2))
}

func uniq(arr []int) (uniq []int) {
	m := make(map[int]bool)

	for _, ele := range arr {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}
	return
}

func findDifference(nums1 []int, nums2 []int) (ans [][]int) {
	nums1, nums2 = uniq(nums1), uniq(nums2)
	tmp := []int{}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				tmp = append(tmp, nums1[i])
			}
		}
	}
	for _, t := range tmp {
		for i := 0; i < len(nums1); i++ {
			if nums1[i] == t {
				nums1 = append(nums1[:i], nums1[i+1:]...)
			}
		}
		for i := 0; i < len(nums2); i++ {
			if nums2[i] == t {
				nums2 = append(nums2[:i], nums2[i+1:]...)
			}
		}
	}
	ans = append(ans, nums1)
	ans = append(ans, nums2)
	return
}
