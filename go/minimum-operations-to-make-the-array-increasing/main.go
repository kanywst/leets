package main

import "fmt"

func main() {
	nums := []int{1, 1, 1}
	// nums := []int{1, 5, 2, 4, 1}
	// n,nums[i] 5,2
	//1,5,6-2=4,7-4=3,8-1=7
	// nums[i] = 2
	// n = 6
	// cnt = n-nums[i]
	fmt.Println(minOperations(nums))
}

func minOperations(nums []int) (cnt int) {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	for i := 0; i < len(tmp)-1; i++ {
		if tmp[i] >= tmp[i+1] {
			// nums[i+1] = nums[i] + 1
			tmp[i+1] = tmp[i] + 1
		}
	}
	for i := 0; i < len(nums); i++ {
		cnt += tmp[i] - nums[i]
	}
	return
}
