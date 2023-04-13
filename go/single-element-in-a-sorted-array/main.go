package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	fmt.Println(singleNonDuplicate(nums))
}

func singleNonDuplicate(nums []int) (ans int) {
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
