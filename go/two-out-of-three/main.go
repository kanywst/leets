package main

import "fmt"

func main() {
	nums1, nums2, nums3 := []int{1, 1, 3, 2}, []int{2, 3}, []int{3}
	fmt.Println(twoOutOfThree(nums1, nums2, nums3))
}

func uniq(arr []int) []int {
	m := make(map[int]bool)
	uniq := []int{}

	for _, ele := range arr {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}
	return uniq
}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) (ans []int) {
	m := make(map[int]int)
	for _, n := range uniq(nums1) {
		m[n] += 1
	}
	for _, n := range uniq(nums2) {
		m[n] += 1
	}
	for _, n := range uniq(nums3) {
		m[n] += 1
	}
	for k, v := range m {
		if v >= 2 {
			ans = append(ans, k)
		}
	}
	return
}
