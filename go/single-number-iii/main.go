package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) (ans []int) {
	m := make(map[int]int)
	for _, i := range nums {
		if _, ok := m[i]; ok {
			m[i] += 1
		} else {
			m[i] = 0
		}
	}
	for k, v := range m {
		if v == 0 {
			ans = append(ans, k)
		}
	}
	return
}
