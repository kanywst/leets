package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	// nums := []int{1, 1}
	fmt.Println(findDisappearedNumbers(nums))
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

func findDisappearedNumbers(nums []int) (ans []int) {
	tmp := []int{}
	for i := 1; i <= len(nums); i++ {
		tmp = append(tmp, i)
	}
	nums = uniq(nums)
	m := make(map[int]bool)
	for _, n := range nums {
		m[n] = true
	}
	for _, i := range tmp {
		if _, ok := m[i]; !ok {
			ans = append(ans, i)
		}
	}
	return
}
