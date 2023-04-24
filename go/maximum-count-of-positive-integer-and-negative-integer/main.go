package main

import "fmt"

func main() {
	nums := []int{-2, -1, -1, 1, 2, 3}
	fmt.Println(maximumCount(nums))
}

func maximumCount(nums []int) int {
	p, n := 0, 0
	for _, i := range nums {
		if i > 0 {
			p += 1
		}
		if i < 0 {
			n += 1
		}
	}
	if p >= n {
		return p
	}
	return n
}
