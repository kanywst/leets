package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func max(nums []int) int {
	ans := -999999
	for i := 0; i < len(nums); i++ {
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans
}

func maxProfit(prices []int) int {
	min := math.MaxUint32
	res := 0
	for _, price := range prices {
		if price > min {
			if price-min > res {
				res = price - min
			}
		} else {
			min = price
		}
	}
	return res
}
