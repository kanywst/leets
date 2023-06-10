package main

import "fmt"

func main() {
	nums1, nums2 := []int{1, 2, 2, 1}, []int{2, 2}
	fmt.Println(intersect(nums1, nums2))
}

func makeMapInt(nums1 []int) map[int]int {
	m1 := make(map[int]int)
	for _, i := range nums1 {
		m1[i] += 1
	}
	return m1
}

func appendAns(ans []int, k int, v int) []int {
	for i := 0; i < v; i++ {
		ans = append(ans, k)
	}
	return ans
}

func intersect(nums1 []int, nums2 []int) (ans []int) {
	m1, m2 := makeMapInt(nums1), makeMapInt(nums2)
	for k, v := range m1 {
		if v <= m2[k] {
			ans = appendAns(ans, k, v)
			continue
		}
		ans = appendAns(ans, k, m2[k])
	}
	return
}
