package main

import "fmt"

func main() {
	nums := []int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}
	fmt.Println(arraySign(nums))
}

func arraySign(nums []int) (ans int) {
	cnt := 0
	for _, i := range nums {
		if i == 0 {
			ans = 0
			return
		} else if i > 0 {
			cnt += 1
		}
	}
	if (len(nums)-cnt)%2 == 0 {
		ans = 1
	} else {
		ans = -1
	}
	return
}
