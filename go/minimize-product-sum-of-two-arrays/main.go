package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1, nums2 := []int{5, 3, 4, 2}, []int{4, 2, 2, 5}
	fmt.Println(minProductSum(nums1, nums2))
}

func reverse(array []int) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}

func minProductSum(nums1 []int, nums2 []int) (ans int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	reverse(nums2)
	for i := 0; i < len(nums1); i++ {
		ans += nums1[i] * nums2[i]
	}
	return
}
