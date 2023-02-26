package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	k := len(nums)
	if k == 0 {
		pattern := []int{}
		return [][]int{pattern}
	}

	ans := [][]int{}
	for i := range nums {
		aRest := make([]int, len(nums)-1)
		for j := 0; j < i; j++ {
			aRest[j] = nums[j]
		}
		for j := i + 1; j < len(nums); j++ {
			aRest[j-1] = nums[j]
		}

		childPatterns := permute(aRest)
		for _, childPattern := range childPatterns {
			pattern := append([]int{nums[i]}, childPattern...)
			ans = append(ans, pattern)
		}
	}

	return ans
}
