package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{10, 4, 8, 3}
	fmt.Println(leftRigthDifference(nums))
}

func sum(nums []int) (sum int) {
	for _, x := range nums {
		sum += x
	}
	return
}

func leftRigthDifference(nums []int) (ans []int) {
	leftSum := make([]int, len(nums))
	rightSum := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		tmp := nums[:i]
		leftSum[i] = sum(tmp)
	}
	for i := 0; i < len(nums); i++ {
		tmp := nums[i+1:]
		rightSum[i] = sum(tmp)
	}
	for i := 0; i < len(nums); i++ {
		ans = append(ans, int(math.Abs(float64(rightSum[i])-float64(leftSum[i]))))
	}
	return
}
