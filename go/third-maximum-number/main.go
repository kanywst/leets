package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(thirdMax(nums))
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

func thirdMax(nums []int) int {
	nn := uniq(nums)
	sort.Ints(nn)
	if len(nn) == 1 {
		return nn[0]
	}
	if len(nn) == 2 {
		return nn[1]
	}
	return nn[len(nn)-3]
}
