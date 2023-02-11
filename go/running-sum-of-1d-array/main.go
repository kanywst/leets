package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(runningSum(nums))
}

func runningSum(nums []int) (answer []int) {
	tmp := 0
	answer = make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		tmp += nums[i]
		answer[i] = tmp
	}
	return
}
