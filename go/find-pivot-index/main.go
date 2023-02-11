package main

import "fmt"

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex((nums)))

}

func pivotIndex(nums []int) (answer int) {
	answer = -1
	leftSum := 0
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	for i := 0; i < len(nums); i++ {
		leftSum += nums[i]
		if leftSum-nums[i] == sum-leftSum {
			answer = i
		}
	}
	return
}
