package main

import "fmt"

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}

func findDuplicate(nums []int) (ans int) {
	m := make(map[int]int)
	for _, i := range nums {
		m[i] += 1
	}
	for k, v := range m {
		if v >= 2 {
			ans = k
		}
	}
	return
}
