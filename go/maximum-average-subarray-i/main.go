package main

import (
	"fmt"
)

func main() {
	nums, k := []int{1, 12, -5, -6, 50, 3}, 4
	fmt.Println(findMaxAverage(nums, k))
}

func sum(x []int) (ans int) {
	for _, xx := range x {
		ans += xx
	}
	return
}

func findMaxAverage(nums []int, k int) float64 {
	tmp := 0.0
	for i := 0; i < len(nums)-k+1; i++ {
		t := float64(sum(nums[i:i+k])) / float64(k)
		if tmp == 0 {
			tmp = t
			continue
		}
		if t > tmp {
			tmp = t
		}
	}
	return tmp
}
