package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) (ans int) {
	m, cnt, ans := make(map[int]int), 0, 0
	for _, i := range nums {
		m[i] += 1
	}
	for i, j := range m {
		if cnt < j {
			cnt, ans = j, i
		}
	}
	return
}
