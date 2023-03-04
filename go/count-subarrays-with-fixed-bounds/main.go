package main

import "fmt"

func main() {
	nums, minK, maxK := []int{1, 3, 5, 2, 7, 5}, 1, 5
	// nums, minK, maxK := []int{1, 1, 1, 1}, 1, 1
	fmt.Println(countSubarrays(nums, minK, maxK))
}

func r_max(x, y int) (ans int) {
	if x > y {
		ans = x
	} else {
		ans = y
	}
	return
}

func r_min(x, y int) (ans int) {
	if x < y {
		ans = x
	} else {
		ans = y
	}
	return
}

func countSubarrays(nums []int, minK int, maxK int) int64 {
	start, min, max, res := 0, -1, -1, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < minK || nums[i] > maxK {
			min, max = -1, -1
			start = i + 1
		}
		if nums[i] == minK {
			min = i
		}

		if nums[i] == maxK {
			max = i
		}
		res += r_max(0, r_min(min, max)-start+1)
	}
	return int64(res)
}
