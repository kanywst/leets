package main

import "fmt"

func main() {
	nums, diff := []int{0, 1, 4, 6, 7, 10}, 3
	fmt.Println(arithmeticTriplets(nums, diff))
}

func arithmeticTriplets(nums []int, diff int) (cnt int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[j]-nums[i] == diff && nums[k]-nums[j] == diff {
					cnt += 1
				}
			}
		}
	}
	return
}
