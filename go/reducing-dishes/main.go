package main

import (
	"fmt"
	"sort"
)

func main() {
	// satisfaction := []int{-1, -8, 0, 5, -9}
	// satisfaction := []int{4, 3, 2}
	satisfaction := []int{-1, -4, -5}
	fmt.Println(maxSatisfaction(satisfaction))
}

func sum(nums []int) (ans int) {
	ii := 1
	for i := 0; i < len(nums); i++ {
		fmt.Println("***", ii, nums[i])
		ans += ii * nums[i]
		ii++
	}
	return
}

func maxSatisfaction(satisfaction []int) (max int) {
	sort.Ints(satisfaction)
	for i := 0; i < len(satisfaction); i++ {
		tmp := sum(satisfaction[i:])
		fmt.Println(satisfaction[i:], tmp)
		if tmp > max {
			max = tmp
		}
	}
	return
}
