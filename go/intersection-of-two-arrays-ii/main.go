package main

import "fmt"

func main() {
	nums1, nums2 := []int{1, 2, 2, 1}, []int{2, 2}
	fmt.Println(intersect(nums1, nums2))
}

func intersect(nums1 []int, nums2 []int) (ans []int) {
	m1, m2 := make(map[int]int), make(map[int]int)
	for _, i := range nums1 {
		m1[i] += 1
	}
	for _, i := range nums2 {
		m2[i] += 1
	}
	for k, v := range m1 {
		if v <= m2[k] {
			for i := 0; i < v; i++ {
				ans = append(ans, k)
			}
		}
		if v > m2[k] {
			for i := 0; i < m2[k]; i++ {
				ans = append(ans, k)
			}
		}
	}
	return
}
