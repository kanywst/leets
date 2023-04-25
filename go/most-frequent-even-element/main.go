package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{0, 1, 2, 2, 4, 4, 1}
	fmt.Println(mostFrequentEven(nums))
}

func mostFrequentEven(nums []int) int {
	m := make(map[int]int)
	for _, n := range nums {
		if n%2 == 0 {
			m[n] += 1
		}
	}
	most := -1
	for _, v := range m {
		if v > most {
			most = v
		}
	}
	ans := math.MaxInt
	for k, v := range m {
		if v == most {
			if k < ans {
				ans = k
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
