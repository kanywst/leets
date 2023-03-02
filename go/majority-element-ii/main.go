package main

import "fmt"

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) (ans []int) {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
	}
	for k, v := range m {
		if v > len(nums)/3 {
			ans = append(ans, k)
		}
	}
	return
}
