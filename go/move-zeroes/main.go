package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	sort.Ints(nums)
	c := 0
	for _, j := range nums {
		if j != 0 {
			break
		} else {
			c += 1
		}
	}
}
