package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{1, 2, 0}
	// nums := []int{3, 4, -1, 1}
	// nums := []int{7, 8, 9, 11, 12}
	// nums := []int{-1, -2}
	nums := []int{0, 2, 2, 1, 1}
	fmt.Println(firstMissingPositive(nums))
}

func firstMissingPositive(nums []int) (ans int) {
	sort.Ints(nums)
	ii, j := 1, 0
	for {
		if len(nums) == 0 || nums[j] > 0 {
			break
		} else if nums[j] <= 0 {
			nums = nums[j+1:]
		} else {
			j++
		}
	}
	fmt.Println("nums", nums)

	m := make(map[int]bool)
	uniq := []int{}

	for _, ele := range nums {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}

	fmt.Println(uniq)

	for i := 0; i < len(uniq); i++ {
		if len(uniq) == 0 {
			break
		}
		if uniq[i] > 0 {
			fmt.Println(uniq[i], ii)
			if uniq[i] != ii {
				break
			}
		}
		ii++
	}
	ans = ii
	return
}
