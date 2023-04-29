package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := "abc1111"
	fmt.Println(secondHighest(s))
}

func secondHighest(s string) int {
	tmp := []int{}
	for _, i := range strings.Split(s, "") {
		n, e := strconv.Atoi(i)
		if e == nil {
			tmp = append(tmp, n)
		}
	}
	m := make(map[int]int)
	for _, i := range tmp {
		m[i] += 1
	}
	nums := []int{}
	for k, _ := range m {
		nums = append(nums, k)
	}
	sort.Ints(nums)
	if len(nums) < 2 {
		return -1
	}
	return nums[len(nums)-2]
}
