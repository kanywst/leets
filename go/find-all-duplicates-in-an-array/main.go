package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums))
}

func findDuplicates(nums []int) (ans []int) {
	m := make(map[int]int)
	for _, n := range nums {
		m[n] += 1
	}
	for k, v := range m {
		if v >= 2 {
			ans = append(ans, k)
		}
	}
	sort.Ints(ans)
	return
}
