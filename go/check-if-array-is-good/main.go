package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 1, 3}
	fmt.Println(isGood(nums))
}

// 1 から n - 1 をちょうど1回含み
// n が2回出現する
// 長さ n + 1 の配列である

func isGood(nums []int) bool {
	sort.Ints(nums)
	n := nums[len(nums)-1]
	if len(nums) != n+1 {
		return false
	}
	m := make(map[int]int)
	for _, nn := range nums {
		m[nn] += 1
	}
	for k, v := range m {
		if k != n {
			if v != 1 {
				return false
			}
		}
		if k == n {
			if m[k] != 2 {
				return false
			}
		}
	}
	return true
}
