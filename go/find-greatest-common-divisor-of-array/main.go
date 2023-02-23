package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 5, 6, 9, 10}
	fmt.Println(findGCD(nums))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func findGCD(nums []int) int {
	sort.Ints(nums)
	min := nums[0]
	max := nums[len(nums)-1]

	return gcd(min, max)
}
