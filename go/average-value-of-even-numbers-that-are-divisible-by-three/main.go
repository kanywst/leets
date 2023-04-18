package main

import "fmt"

func main() {
	nums := []int{1, 3, 6, 10, 12, 15}
	fmt.Println(averageValue(nums))
}

func averageValue(nums []int) int {
	tmp := []int{}
	for _, i := range nums {
		if i%2 == 0 && i%3 == 0 {
			tmp = append(tmp, i)
		}
	}
	sum := 0
	for _, i := range tmp {
		sum += i
	}
	if len(tmp) == 0 {
		return 0
	}
	return sum / len(tmp)
}
