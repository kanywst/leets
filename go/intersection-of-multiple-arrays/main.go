package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{
		[]int{3, 1, 2, 4, 5},
		[]int{1, 2, 3, 4},
		[]int{3, 4, 5, 6},
	}
	fmt.Println(intersection(nums))
}

func intersection(nums [][]int) (ans []int) {
	m := make(map[int]int)
	for _, n := range nums {
		for _, i := range n {
			m[i] += 1
			if m[i] == len(nums) {
				ans = append(ans, i)
			}
		}
	}
	sort.Ints(ans)
	return
}
