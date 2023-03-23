package main

import "fmt"

func main() {
	nums, target := []int{4, 5, 6, 7, 0, 1, 2}, 0
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	for i, j := range nums {
		if j == target {
			return i
		}
	}
	return -1
}
