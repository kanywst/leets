package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums, target := []int{1, 3, 5, 6}, 5
	nums, target := []int{1, 3, 5, 6}, 2
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) (ans int) {
	tmp := append(nums, []int{target}...)
	sort.Ints(tmp)
	for i := 0; i < len(tmp); i++ {
		if tmp[i] == target {
			ans = i
			break
		}
	}
	return
}
