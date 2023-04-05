package main

import "fmt"

func main() {
	nums1, nums2 := []int{1, 2, 2, 1}, []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) (ans []int) {
	m := make(map[int]int)
	for _, i := range nums1 {
		for _, j := range nums2 {
			if i == j {
				if _, ok := m[i]; !ok {
					m[i] = j
					break
				}
			}
		}
	}
	for _, i := range m {
		ans = append(ans, i)
	}
	return
}
