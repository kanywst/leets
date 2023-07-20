package main

import "fmt"

func main() {
	nums := []int{2, 7, 1, 19, 18, 3}
	fmt.Println(sumOfSquares(nums))
}

func sumOfSquares(nums []int) (ans int) {
	for i := 0; i < len(nums); i++ {
		if len(nums)%(i+1) == 0 {
			ans += nums[i] * nums[i]
		}
	}
	return
}
