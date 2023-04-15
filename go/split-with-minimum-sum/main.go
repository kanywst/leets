package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	num := 4325
	fmt.Println(splitNum(num))
}

func splitNum(num int) int {
	s := strings.Split(strconv.Itoa(num), "")
	nums := []int{}
	for _, i := range s {
		n, _ := strconv.Atoi(i)
		nums = append(nums, n)
	}
	sort.Ints(nums)
	n1 := len(nums) / 2
	//n2 := len(nums) - n1
	nums1 := nums[:n1]
	nums2 := nums[n1:]
	fmt.Println(nums1, nums2)
	return 0
}
