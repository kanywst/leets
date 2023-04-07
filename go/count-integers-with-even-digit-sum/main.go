package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num := 4
	fmt.Println(countEven(num))
}

func sum(nums []int) (ans int) {
	for _, i := range nums {
		ans += i
	}
	return
}

func countEven(num int) (cnt int) {
	for i := 1; i <= num; i++ {
		nums := []int{}
		for _, i := range strings.Split(strconv.Itoa(i), "") {
			n, _ := strconv.Atoi(i)
			nums = append(nums, n)
		}
		if sum(nums)%2 == 0 {
			cnt += 1
		}
	}
	return
}
