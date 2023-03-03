package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))
}

func summaryRanges(nums []int) (ans []string) {
	nums, ans, n_l := append(nums, -1), []string{}, []int{}
	for i := 0; i < len(nums)-1; i++ {
		n_l = append(n_l, nums[i])
		if nums[i]+1 == nums[i+1] {
		} else {
			if len(n_l) > 1 {
				ans = append(ans, strconv.Itoa(n_l[0])+"->"+strconv.Itoa(n_l[len(n_l)-1]))
			} else {
				ans = append(ans, strconv.Itoa(n_l[0]))
			}
			n_l = []int{}
		}
	}
	return
}
