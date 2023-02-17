package main

import "fmt"

func main() {
	nums1, nums2 := []int{4, 1, 2}, []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	for i, j := range nums1 {
		tmp := []int{}
		for k, l := range nums2 {
			if j == l {
				tmp = nums2[k:]
				break
			}
		}
		for _, l := range tmp {
			if j < l {
				ans[i] = l
				break
			} else {
				ans[i] = -1
			}
		}
	}
	return ans
}
