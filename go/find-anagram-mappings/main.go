package main

import "fmt"

func main() {
	nums1, nums2 := []int{12, 28, 46, 32, 50}, []int{50, 12, 32, 46, 28}
	fmt.Println(anagramMappings(nums1, nums2))
}

func anagramMappings(nums1 []int, nums2 []int) (ans []int) {
	m := make(map[int]int)
	for i, j := range nums2 {
		m[j] = i
	}
	for _, n := range nums1 {
		ans = append(ans, m[n])
	}
	return
}
