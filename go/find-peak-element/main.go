package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(findPeakElement(nums))
}

func max(x []int) (ans int) {
	tmp := x[0]
	for i := 1; i < len(x); i++ {
		if tmp < x[i] {
			tmp = x[i]
		}
	}
	ans = tmp
	return
}

func findPeakElement(nums []int) int {
	m, ans := max(nums), -1
	for i := 1; i < len(nums); i++ {
		if nums[i] == m {
			ans = i - 1
			return nums[ans]
		}
	}
	return nums[ans]
}
