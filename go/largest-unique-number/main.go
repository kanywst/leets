package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 7, 3, 9, 4, 9, 8, 3, 1}
	fmt.Println(largestUniqueNumber(nums))
}

func largestUniqueNumber(nums []int) int {
	m := make(map[int]int)
	for _, i := range nums {
		m[i] += 1
	}
	tmp := []int{}
	for k, v := range m {
		if v == 1 {
			tmp = append(tmp, k)
		}
	}
	sort.Ints(tmp)
	if len(tmp) == 0 {
		return -1
	}
	return tmp[len(tmp)-1]
}
