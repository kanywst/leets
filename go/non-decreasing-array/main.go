package main

import "fmt"

func main() {
	nums := []int{4, 2, 3}
	fmt.Println(checkPossibility(nums))
}

func checkPossibility(nums []int) (ans bool) {
	cnt := 0
	ans = true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			cnt += 1
		} else if cnt > 0 {
			ans = false
			break
		}
	}
	return
}
