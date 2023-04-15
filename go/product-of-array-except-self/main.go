package main

import "fmt"

func main() {
	nums := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) (ans []int) {
	ans = make([]int, len(nums))
	tmp_num := 1
	for i, j := range nums {
		fmt.Println(ans, tmp_num, j)
		ans[i] *= j
		tmp_num *= ans[i]
	}
	fmt.Println(ans, tmp_num)
	return
}
