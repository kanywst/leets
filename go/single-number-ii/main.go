package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) (ans int) {
	m := make(map[int]int)
	for _, i := range nums {
		m[i] += 1
	}
	for k, v := range m {
		if v == 1 {
			ans = k
			break
		}
	}
	return
}
