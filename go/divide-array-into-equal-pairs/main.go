package main

import "fmt"

func main() {
	// nums := []int{3, 2, 3, 2, 2, 2}
	nums := []int{1, 2, 3, 4}
	fmt.Println(divideArray(nums))
}

func divideArray(nums []int) bool {
	m := make(map[int]int)
	for _, n := range nums {
		m[n] += 1
	}
	for _, v := range m {
		if v%2 != 0 {
			return false
		}
	}
	return true
}
